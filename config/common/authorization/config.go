package authorization

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_authorization_project_custom_role", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_authorization_project_role_assignment", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_authorization_folder_role_assignment", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_authorization_organization_role_assignment", func(r *config.Resource) {
	})
}
