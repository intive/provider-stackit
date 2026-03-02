package objectstorage

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_objectstorage_bucket", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_objectstorage_credential", func(r *config.Resource) {
		r.References["credentials_group_id"] = config.Reference{
			TerraformName: "stackit_objectstorage_credentials_group",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("credentials_group_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_objectstorage_credentials_group", func(r *config.Resource) {
	})
}
