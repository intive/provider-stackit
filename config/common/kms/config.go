package kms

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_kms_keyring", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_kms_key", func(r *config.Resource) {
		r.References["keyring_id"] = config.Reference{
			TerraformName: "stackit_kms_keyring",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("keyring_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_kms_wrapping_key", func(r *config.Resource) {
		r.References["keyring_id"] = config.Reference{
			TerraformName: "stackit_kms_keyring",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("keyring_id",true)`,
		}
	})
}
