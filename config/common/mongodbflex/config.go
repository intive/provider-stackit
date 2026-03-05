package mongodbflex

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_mongodbflex_user", func(r *config.Resource) {
		r.ShortGroup = "mongodbflex"
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_mongodbflex_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
