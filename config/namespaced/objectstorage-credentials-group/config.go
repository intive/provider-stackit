package objectstoragecredentialsgroup

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_objectstorage_credentials_group", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})
}
