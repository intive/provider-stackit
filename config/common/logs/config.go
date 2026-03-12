package logs

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_logs_instance", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_logs_access_token", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_logs_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
