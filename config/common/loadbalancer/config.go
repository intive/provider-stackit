package loadbalancer

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_loadbalancer", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["networks.network_id"] = config.Reference{
			TerraformName: "stackit_network",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_loadbalancer_observability_credential", func(r *config.Resource) {
	})
}
