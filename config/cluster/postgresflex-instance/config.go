package postgresflexinstance

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_instance", func(r *config.Resource) {
		r.ShortGroup = "postgresflex"
	})
}
