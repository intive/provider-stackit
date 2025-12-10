package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	objectstorageBucketCluster "github.com/intive/provider-stackit/config/cluster/objectstorage-bucket"
	objectstorageCredentialsCluster "github.com/intive/provider-stackit/config/cluster/objectstorage-credential"
	objectstorageCredentialsGroupCluster "github.com/intive/provider-stackit/config/cluster/objectstorage-credentials-group"
	postgresflexDatabaseCluster "github.com/intive/provider-stackit/config/cluster/postgresflex-database"
	postgresflexInstanceCluster "github.com/intive/provider-stackit/config/cluster/postgresflex-instance"
	postgresflexUserCluster "github.com/intive/provider-stackit/config/cluster/postgresflex-user"
	rediscredentialCluster "github.com/intive/provider-stackit/config/cluster/redis-credential"
	redisInstanceCluster "github.com/intive/provider-stackit/config/cluster/redis-instance"
	objectstorageBucketNamespaced "github.com/intive/provider-stackit/config/namespaced/objectstorage-bucket"
	objectstorageCredentialsNamespaced "github.com/intive/provider-stackit/config/namespaced/objectstorage-credential"
	objectstorageCredentialsGroupNamespaced "github.com/intive/provider-stackit/config/namespaced/objectstorage-credentials-group"
	postgresflexDatabaseNamespaced "github.com/intive/provider-stackit/config/namespaced/postgresflex-database"
	postgresflexInstanceNamespaced "github.com/intive/provider-stackit/config/namespaced/postgresflex-instance"
	postgresflexUserNamespaced "github.com/intive/provider-stackit/config/namespaced/postgresflex-user"
	rediscredentialNamespaced "github.com/intive/provider-stackit/config/namespaced/redis-credential"
	redisInstanceNamespaced "github.com/intive/provider-stackit/config/namespaced/redis-instance"
)

const (
	resourcePrefix = "stackit"
	modulePath     = "github.com/intive/provider-stackit"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorageBucketCluster.Configure,
		redisInstanceCluster.Configure,
		rediscredentialCluster.Configure,
		postgresflexInstanceCluster.Configure,
		postgresflexDatabaseCluster.Configure,
		postgresflexUserCluster.Configure,
		objectstorageCredentialsGroupCluster.Configure,
		objectstorageCredentialsCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorageBucketNamespaced.Configure,
		redisInstanceNamespaced.Configure,
		rediscredentialNamespaced.Configure,
		postgresflexInstanceNamespaced.Configure,
		postgresflexDatabaseNamespaced.Configure,
		postgresflexUserNamespaced.Configure,
		objectstorageCredentialsGroupNamespaced.Configure,
		objectstorageCredentialsNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
