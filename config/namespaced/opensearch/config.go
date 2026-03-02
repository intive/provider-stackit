package opensearch

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_opensearch_instance", func(r *config.Resource) {
		r.ShortGroup = "opensearch"
	})
	p.AddResourceConfigurator("stackit_opensearch_credential", func(r *config.Resource) {
		r.ShortGroup = "opensearch"
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_opensearch_instance",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
		}
	})
}
