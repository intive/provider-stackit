package git

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_git", func(r *config.Resource) {
		r.ShortGroup = "git"
	})
}
