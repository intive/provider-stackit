# Provider STACKIT

`provider-stackit` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) codegeneration tools and exposes XRM-conformant managed resources for the [STACKIT](https://www.stackit.de) API.

## Usage

### Installation

To install the provider, use the following resource definition:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-stackit
  namespace: crossplane-system
spec:
  package: xpkg.upbound.io/intive/provider-stackit:v0.0.1
```

This will install the provider in the `crossplane-system` namespace and install CRDs and controllers for the provider.

### Configuration

To configure the provider using a [service account key](https://registry.terraform.io/providers/stackitcloud/stackit/latest/docs#key-flow), use the following resource definition:

## Resource Status

| Service | Resource | Status | Notes |
|---------|----------|--------|-------|
| authorization | OrganizationRoleAssignment | blocked | requires `iam` experiment (fix #5) |
| authorization | ProjectRoleAssignment | blocked | requires `iam` experiment (fix #5) |
| cdn | CustomDomain | beta | requires `enable_beta_resources` (fix #4) |
| cdn | Distribution | beta | requires `enable_beta_resources` (fix #4) |
| compute | AffinityGroup | ✅ | |
| compute | Image | blocked | requires uploading a local qcow2 image file (`localFilePath`), not testable in CI |
| compute | KeyPair | ✅ | |
| compute | Server | ✅ | |
| compute | Volume | ✅ | |
| compute | NetworkInterfaceAttach | ✅ | |
| compute | VolumeAttach | ✅ | |
| dns | RecordSet | ✅ | |
| dns | Zone | ✅ | |
| edgecloud | Instance | beta | requires `enable_beta_resources` (fix #4) |
| git | Git | beta | requires `enable_beta_resources` (fix #4) |
| kms | Key | ✅ | |
| kms | Keyring | ✅ | |
| kms | WrappingKey | ✅ | |
| loadbalancer | Loadbalancer | ✅ | |
| logme | Credential | ✅ | |
| logme | Instance | ✅ | |
| logs | AccessToken | ✅ | |
| logs | Instance | ✅ | |
| mariadb | Credential | ✅ | |
| mariadb | Instance | ✅ | |
| modelserving | Token | beta | requires `enable_beta_resources` (fix #4) |
| mongodbflex | Instance | ✅ | |
| mongodbflex | User | ✅ | |
| network | Area | blocked | SA lacks org-level network area permissions (403 Forbidden) |
| network | AreaRegion | blocked | depends on Area |
| network | AreaRoute | blocked | depends on Area |
| network | Network | ✅ | |
| network | NetworkInterface | ✅ | |
| network | NetworkInterfaceAttach | moved | resource moved to compute group |
| objectstorage | Bucket | ✅ | |
| objectstorage | CredentialsGroup | ✅ | |
| observability | AlertGroup | ✅ | |
| observability | Credential | ✅ | |
| observability | Instance | ✅ | |
| opensearch | Credential | ✅ | |
| opensearch | Instance | ✅ | |
| postgresflex | Database | blocked | deletion blocked: API returns 500 for database operations while instance is deleting (fix #6) |
| postgresflex | Instance | ✅ | |
| postgresflex | User | blocked | deletion blocked: API returns 500 for user operations while instance is deleting (fix #6) |
| rabbitmq | Credential | ✅ | |
| rabbitmq | Instance | ✅ | |
| redis | Credential | ✅ | |
| redis | Instance | ✅ | |
| resourcemanager | Folder | blocked | service account lacks org-level folder creation permissions |
| resourcemanager | Project | blocked | service account lacks org-level project creation permissions |
| scf | Organization | beta | requires `enable_beta_resources` (fix #4) |
| secretsmanager | Instance | ✅ | |
| secretsmanager | User | ✅ | |
| serviceaccount | Account | ✅ | |
| serviceaccount | AccountKey | ✅ | |
| sfs | ResourcePool | beta | requires `enable_beta_resources` (fix #4) |
| ske | Cluster | blocked | SKE service not enabled on test project (403 Forbidden) |
| ske | Kubeconfig | blocked | depends on Cluster |
| sqlserverflex | Instance | ✅ | |
| sqlserverflex | User | ✅ | |

### Status Legend

- ✅ **Tested and working**
- **not tested** — not yet covered by uptest
- **blocked** — known upstream issue documented in [fixes.md](fixes.md)
- **beta** — resource requires `enable_beta_resources` support (fix #4 in fixes.md)

---

```yaml
apiVersion: stackit.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: provider-stackit-config
spec:
  credentials:
    source: Secret
    secretRef:
      name: stackit-credentials
      namespace: crossplane-system
      key: credentials
---
apiVersion: v1
kind: Secret
metadata:
  name: stackit-credentials
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "default_region": "eu01",
      "service_account_key": "<JSON-OF-SERVICE-ACCOUNT-KEY"
    }
```
