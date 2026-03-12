package edgecloud

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_edgecloud_instance", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_edgecloud_kubeconfig", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_edgecloud_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_edgecloud_token", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_edgecloud_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
