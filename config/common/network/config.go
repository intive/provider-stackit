package network

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("stackit_network_interface", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			TerraformName: "stackit_network",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_id",true)`,
		}
	})
}
