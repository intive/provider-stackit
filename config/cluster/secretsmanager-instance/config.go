package secretsmanagerinstance

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_secretsmanager_instance", func(r *config.Resource) {
		r.ShortGroup = "secretsmanager"
	})
}
