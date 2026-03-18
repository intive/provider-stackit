package modelserving

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_modelserving_token", func(r *config.Resource) {
	})
}
