package dns

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_dns_zone", func(r *config.Resource) {
	})

	p.AddResourceConfigurator("stackit_dns_record_set", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			TerraformName: "stackit_dns_zone",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("zone_id",true)`,
		}
	})
}
