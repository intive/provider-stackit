package scf

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_scf_organization", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_scf_organization_manager", func(r *config.Resource) {
		r.References["org_id"] = config.Reference{
			TerraformName: "stackit_scf_organization",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("org_id",true)`,
		}
	})
}
