package cdn

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_cdn_distribution", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_cdn_custom_domain", func(r *config.Resource) {
		r.References["distribution_id"] = config.Reference{
			TerraformName: "stackit_cdn_distribution",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("distribution_id",true)`,
		}
	})
}
