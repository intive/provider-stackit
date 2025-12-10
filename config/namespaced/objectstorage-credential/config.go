package objectstoragecredential

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_objectstorage_credential", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["credentials_group_id"] = config.Reference{
			TerraformName: "stackit_objectstorage_credentials_group",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("credentials_group_id",true)`,
		}
	})
}
