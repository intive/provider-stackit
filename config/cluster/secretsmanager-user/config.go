package secretsmanageruser

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_secretsmanager_user", func(r *config.Resource) {
		r.ShortGroup = "secretsmanager"
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_secretsmanager_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
