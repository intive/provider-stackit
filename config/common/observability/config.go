package observability

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_observability_instance", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_observability_credential", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_observability_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_observability_alertgroup", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_observability_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_observability_logalertgroup", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_observability_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_observability_scrapeconfig", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_observability_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
