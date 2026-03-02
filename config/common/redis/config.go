package redis

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_redis_credential", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_redis_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
