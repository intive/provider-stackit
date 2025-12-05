package objectstoragebucket

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_objectstorage_bucket", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})
}
