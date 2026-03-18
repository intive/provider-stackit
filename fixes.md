# Fixes in Upstream Terraform Provider

This file documents issues in the upstream terraform provider
(`github.com/stackitcloud/terraform-provider-stackit`) that affect
this crossplane provider.

---

## 1. Handle 404 (Not Found) gracefully during Delete and Read ✅ DONE

**Affected provider:** `github.com/stackitcloud/terraform-provider-stackit`

### Problem

Several STACKIT resources have parent-child relationships where deleting the
parent in the cloud also cascades and deletes all child resources. In the
crossplane/uptest testing lifecycle, all Kubernetes CRs are deleted
simultaneously. This means:

1. The parent (e.g. `Instance`) CR triggers deletion → cloud deletes the
   instance **and** all its children (users, databases, credentials, etc.)
2. The child CRs (e.g. `User`, `Database`) also trigger deletion → provider
   calls the STACKIT API to delete the child resources → receives **404 Not
   Found** because the cloud already deleted them
3. The provider returns an error on 404 → the crossplane finalizer is never
   removed → the Kubernetes CR is **permanently stuck**

The same issue occurs during reconciliation (`Read`): if an external actor
deletes a resource, the `Read` function should signal "resource no longer
exists" (by setting the Terraform ID to `""`) rather than returning an error.

### Fix Required

In every resource's `Read` / `Delete` function in the terraform provider, a
404 / "not found" API response should be treated as a success:

- **`Read`**: return `d.SetId("")` (resource is gone, remove from state)
- **`Delete`**: return `nil` (resource is already gone, deletion succeeded)

Example pattern:

```go
// In Read:
instance, err := client.GetInstance(ctx, projectId, instanceId)
if err != nil {
    if isNotFound(err) {
        d.SetId("")
        return nil
    }
    return diag.FromErr(err)
}

// In Delete:
err := client.DeleteInstance(ctx, projectId, instanceId)
if err != nil {
    if isNotFound(err) {
        return nil // already deleted
    }
    return diag.FromErr(err)
}
```

### Affected Resource Groups (cascading parent-child deletion)

| Parent | Children (auto-deleted with parent in cloud) |
|--------|----------------------------------------------|
| `stackit_postgresflex_instance` | `stackit_postgresflex_user`, `stackit_postgresflex_database` |
| `stackit_mongodbflex_instance` | `stackit_mongodbflex_user` |
| `stackit_sqlserverflex_instance` | `stackit_sqlserverflex_user` |
| `stackit_logme_instance` | `stackit_logme_credential` |
| `stackit_mariadb_instance` | `stackit_mariadb_credential` |
| `stackit_opensearch_instance` | `stackit_opensearch_credential` |
| `stackit_observability_instance` | `stackit_observability_credential` |
| `stackit_redis_instance` | `stackit_redis_credential` |
| `stackit_rabbitmq_instance` | `stackit_rabbitmq_credential` |
| `stackit_secretsmanager_instance` | `stackit_secretsmanager_user` |
| `stackit_objectstorage_credentials_group` | `stackit_objectstorage_credential` |
| `stackit_dns_zone` | `stackit_dns_record_set` |
| `stackit_kms_keyring` | `stackit_kms_key`, `stackit_kms_wrapping_key` |
| `stackit_ske_cluster` | `stackit_ske_kubeconfig` |
| `stackit_serviceaccount` | `stackit_serviceaccount_key` |
| `stackit_network` | `stackit_network_interface` → `stackit_network_interface_attach` |
| `stackit_network_area` | `stackit_network_area_route`, `stackit_network_area_region` |
| `stackit_server` | `stackit_volume_attach` |

### Note on uptest Deletion Order

`uptest v2` deletes all resources simultaneously (`kubectl delete --wait=false`),
so ordering annotations cannot prevent the race. The only robust fix is 404
handling in the provider itself.

---

## 2. Network resource: ConflictsWith validators triggered during observe ✅ DONE

**Affected provider:** `github.com/stackitcloud/terraform-provider-stackit`

### Problem

`stackit_network`'s `mapFields` function (called during Read/observe) set both
members of several mutually-exclusive attribute pairs simultaneously, causing
`ConfigValidators` to reject the state:

- `ipv4_prefix` **and** `ipv4_prefix_length` were both set to non-null
- `ipv4_prefix_length` **and** `ipv4_gateway` were both set to non-null
- `nameservers` (deprecated) **and** `ipv4_nameservers` were both set to non-null

Result: every observe cycle failed with:
```
Invalid Attribute Combination: These attributes cannot be configured together: [ipv4_prefix,ipv4_prefix_length]
Invalid Attribute Combination: These attributes cannot be configured together: [ipv4_prefix_length,ipv4_gateway]
Error configuring network: You cannot provide both the `nameservers` and `ipv4_nameservers` fields simultaneously.
```

### Fix

In `mapFields` (`stackit/internal/services/iaas/network/resource.go`):

1. **Nameservers**: only update the field the user originally configured. If
   `model.Nameservers` was non-null (user used deprecated field), update only
   `Nameservers` and set `IPv4Nameservers` to null. Otherwise update only
   `IPv4Nameservers` and set `Nameservers` to null.

2. **Prefix**: only set `IPv4PrefixLength` if the user originally had it
   non-null; otherwise set only `IPv4Prefix` (and keep `IPv4PrefixLength`
   null). This preserves the `ipv4_prefix` / `ipv4_prefix_length` choice.

3. **Gateway**: set `IPv4Gateway` to null when `IPv4PrefixLength` is non-null,
   since they conflict.

---

## 3. PostgreSQL Flex (and other database services): 404 during instance creation polling

**Affected package:** `github.com/stackitcloud/stackit-sdk-go/services/postgresflex/wait`
(and equivalent wait packages for other services)

### Problem

In `CreateInstanceWaitHandler`, when `GetInstanceExecute` returns a 404 during
the polling loop, the error is immediately returned as fatal:

```go
s, err := a.GetInstanceExecute(ctx, projectId, region, instanceId)
if err != nil {
    return false, nil, err  // 404 treated as fatal — stops the poller
}
```

The handler already calls `SetSleepBeforeWait(15 * time.Second)` to delay the
first poll, but if the API is still not consistent after 15s, subsequent polls
also fail with 404 and the wait handler aborts. Result:

```
apply failed: Error creating instance: Instance creation waiting:
404 Not Found ... Requested instance with ID: ... cannot be found
```

The same bug exists in `DeleteInstanceWaitHandler`: a 404 should mean
"already deleted, done" but is treated as a fatal error.

### Fix Required

In the SDK's `CreateInstanceWaitHandler` poll function, treat 404 as
"not yet available, keep retrying":

```go
if err != nil {
    oapiErr, ok := err.(*oapierror.GenericOpenAPIError)
    if ok && oapiErr.StatusCode == 404 {
        return false, nil, nil  // keep waiting
    }
    return false, nil, err
}
```

In `DeleteInstanceWaitHandler`, treat 404 as success:

```go
if err != nil {
    oapiErr, ok := err.(*oapierror.GenericOpenAPIError)
    if ok && oapiErr.StatusCode == 404 {
        return true, nil, nil  // deleted, done
    }
    return false, nil, err
}
```

### Affected Services

Same fix needed in the wait packages for all services that use the same pattern:
`postgresflex`, `mongodbflex`, `sqlserverflex`, `logme`, `mariadb`,
`opensearch`, `observability`, `redis`, `rabbitmq`

---

## 4. Beta resources not enabled (git, edgecloud, modelserving, scf)

**Affected:** `stackit_git`, `stackit_edgecloud_instance`, `stackit_modelserving_token`, `stackit_scf_organization`, `stackit_sfs_resource_pool`, `stackit_cdn_distribution`, `stackit_cdn_custom_domain`, `stackit_server_backup_schedule`, `stackit_server_update_schedule`

### Problem

Several STACKIT resources are marked as beta in the Terraform provider and
require an explicit opt-in:

```
The resource "stackit_git" is in beta and beta is not enabled.
To enable it, set STACKIT_TF_ENABLE_BETA_RESOURCES=true or
set "enable_beta_resources" in the provider configuration.
```

The crossplane provider currently has no mechanism to pass
`enable_beta_resources: true` to the embedded Terraform provider.

### Fix Required

Add support for the `enable_beta_resources` provider configuration field
(or the `STACKIT_TF_ENABLE_BETA_RESOURCES` environment variable) in the
crossplane provider's ProviderConfig CRD and in the setup function that
configures the embedded Terraform provider. This will allow users to opt into
beta resources when needed.

---

## 5. Authorization resources require `iam` experiment flag

**Affected:** `stackit_authorization_project_role_assignment`, `stackit_authorization_organization_role_assignment`, `stackit_authorization_project_custom_role`

### Problem

```
stackit_authorization_project_role_assignment is part of the iam experiment,
which is currently disabled by default: Enable the iam experiment by adding
it into your provider block.
```

The crossplane provider has no mechanism to pass `experiments = ["iam"]` to the
embedded Terraform provider configuration.

### Fix Required

Add support for the `experiments` provider configuration field in the
crossplane provider's ProviderConfig CRD and setup function, alongside the
`enable_beta_resources` field from fix #4.

---

## 6. PostgresFlex user/database: API returns 500 during deletion when instance is deleting

**Affected provider:** `github.com/stackitcloud/terraform-provider-stackit`

### Problem

When a `stackit_postgresflex_instance` is deleted and the controller then tries
to delete child `stackit_postgresflex_user` or `stackit_postgresflex_database`
resources (or observe them during the deletion reconcile loop), the STACKIT API
returns `500 Internal Server Error` with the message:

```
{"message":"instance not available yet, unable to find all necessary objects","code":500,"type":"Read"}
```

The provider treats 500 as a transient error and keeps retrying indefinitely.
Since the Crossplane finalizer is never removed, the child CRs are permanently
stuck in `Deleting` state and require manual finalizer removal.

This differs from fix #1 (404 handling) because the API returns 500, not 404.
The sqlserverflex service has the same issue but the 500 resolves itself after
~4 minutes; the postgresflex 500 never resolves.

### Fix Required

In the `Read` function for `stackit_postgresflex_user` and
`stackit_postgresflex_database`, treat a 500 with message containing
`"instance not available yet"` (or check if the parent instance no longer
exists) as "resource is gone" and return `d.SetId("")`:

```go
if err != nil {
    if isNotFound(err) || isInstanceUnavailable(err) {
        d.SetId("") // parent instance gone, treat user/database as deleted
        return nil
    }
    return diag.FromErr(err)
}
```

### Affected Resources

- `stackit_postgresflex_user`
- `stackit_postgresflex_database`
