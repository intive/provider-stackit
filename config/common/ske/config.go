package ske

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_ske_cluster", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_ske_kubeconfig", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			TerraformName: "stackit_ske_cluster",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",false)`,
		}
	})
}
