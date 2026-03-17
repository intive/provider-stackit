package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	// Common config functions
	mariadb "github.com/intive/provider-stackit/config/common/mariadb"
	mongodbflex "github.com/intive/provider-stackit/config/common/mongodbflex"
	objectstorage "github.com/intive/provider-stackit/config/common/objectstorage"
	opensearch "github.com/intive/provider-stackit/config/common/opensearch"
	postgresflex "github.com/intive/provider-stackit/config/common/postgresflex"
	rabbitmq "github.com/intive/provider-stackit/config/common/rabbitmq"
	redis "github.com/intive/provider-stackit/config/common/redis"
	secretsmanager "github.com/intive/provider-stackit/config/common/secretsmanager"
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

	ConfigureCommon(pc)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom cluster-wide config functions
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

	ConfigureCommon(pc)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom namespaced config functions
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func ConfigureCommon(pc *ujconfig.Provider) {
	for _, configure := range []func(provider *ujconfig.Provider){
		objectstorage.Configure,
		redis.Configure,
		postgresflex.Configure,
		secretsmanager.Configure,
		rabbitmq.Configure,
		opensearch.Configure,
		mariadb.Configure,
		mongodbflex.Configure,
	} {
		configure(pc)
	}
}
