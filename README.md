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
