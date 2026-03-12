package loadbalancer

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_loadbalancer", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
	})

	p.AddResourceConfigurator("stackit_loadbalancer_observability_credential", func(r *config.Resource) {
	})
}
