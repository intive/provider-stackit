package sfs

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_sfs_resource_pool", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_sfs_share", func(r *config.Resource) {
		r.References["resource_pool_id"] = config.Reference{
			TerraformName: "stackit_sfs_resource_pool",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("resource_pool_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_sfs_export_policy", func(r *config.Resource) {
	})
}
