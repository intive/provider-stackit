package resourcemanager

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_resourcemanager_folder", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_resourcemanager_project", func(r *config.Resource) {
	})
}
