package postgresflexdatabase

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_database", func(r *config.Resource) {
		r.ShortGroup = "postgresflex"
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
