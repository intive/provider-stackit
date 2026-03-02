package rabbitmqinstance

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_rabbitmq_instance", func(r *config.Resource) {
		r.ShortGroup = "rabbitmq"
	})
}
