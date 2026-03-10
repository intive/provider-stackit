package observability

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_observability_credential", addInstanceReference)
	p.AddResourceConfigurator("stackit_observability_alertgroup", addInstanceReference)
	p.AddResourceConfigurator("stackit_observability_logalertgroup", addInstanceReference)
	p.AddResourceConfigurator("stackit_observability_scrapeconfig", addInstanceReference)
}

func addInstanceReference(r *config.Resource) {
	r.References["instance_id"] = config.Reference{
		TerraformName: "stackit_observability_instance",
		Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("instance_id",true)`,
	}
}
