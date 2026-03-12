package serviceaccount

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_service_account", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
	})

	p.AddResourceConfigurator("stackit_service_account_access_token", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
		r.References["service_account_email"] = config.Reference{
			TerraformName: "stackit_service_account",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("email",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_service_account_key", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
		r.References["service_account_email"] = config.Reference{
			TerraformName: "stackit_service_account",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("email",true)`,
		}
	})
}
