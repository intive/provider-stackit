package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	// Common config functions
	authorization "github.com/intive/provider-stackit/config/common/authorization"
	cdn "github.com/intive/provider-stackit/config/common/cdn"
	compute "github.com/intive/provider-stackit/config/common/compute"
	dns "github.com/intive/provider-stackit/config/common/dns"
	edgecloud "github.com/intive/provider-stackit/config/common/edgecloud"
	git "github.com/intive/provider-stackit/config/common/git"
	kms "github.com/intive/provider-stackit/config/common/kms"
	loadbalancer "github.com/intive/provider-stackit/config/common/loadbalancer"
	logme "github.com/intive/provider-stackit/config/common/logme"
	logs "github.com/intive/provider-stackit/config/common/logs"
	mariadb "github.com/intive/provider-stackit/config/common/mariadb"
	modelserving "github.com/intive/provider-stackit/config/common/modelserving"
	mongodbflex "github.com/intive/provider-stackit/config/common/mongodbflex"
	network "github.com/intive/provider-stackit/config/common/network"
	objectstorage "github.com/intive/provider-stackit/config/common/objectstorage"
	observability "github.com/intive/provider-stackit/config/common/observability"
	opensearch "github.com/intive/provider-stackit/config/common/opensearch"
	postgresflex "github.com/intive/provider-stackit/config/common/postgresflex"
	rabbitmq "github.com/intive/provider-stackit/config/common/rabbitmq"
	redis "github.com/intive/provider-stackit/config/common/redis"
	resourcemanager "github.com/intive/provider-stackit/config/common/resourcemanager"
	scf "github.com/intive/provider-stackit/config/common/scf"
	secretsmanager "github.com/intive/provider-stackit/config/common/secretsmanager"
	serviceaccount "github.com/intive/provider-stackit/config/common/serviceaccount"
	sfs "github.com/intive/provider-stackit/config/common/sfs"
	ske "github.com/intive/provider-stackit/config/common/ske"
	sqlserverflex "github.com/intive/provider-stackit/config/common/sqlserverflex"
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
		network.Configure,
		compute.Configure,
		dns.Configure,
		ske.Configure,
		edgecloud.Configure,
		kms.Configure,
		authorization.Configure,
		resourcemanager.Configure,
		serviceaccount.Configure,
		mariadb.Configure,
		opensearch.Configure,
		mongodbflex.Configure,
		sqlserverflex.Configure,
		logme.Configure,
		logs.Configure,
		observability.Configure,
		cdn.Configure,
		loadbalancer.Configure,
		sfs.Configure,
		git.Configure,
		scf.Configure,
		modelserving.Configure,
	} {
		configure(pc)
	}
}
