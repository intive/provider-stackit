package network

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("stackit_network_area", func(r *config.Resource) {
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("stackit_network_area_region", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_area_id"] = config.Reference{
			TerraformName: "stackit_network_area",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_area_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_network_area_route", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_area_id"] = config.Reference{
			TerraformName: "stackit_network_area",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_area_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_network_interface", func(r *config.Resource) {
		r.ShortGroup = "network"
		// Override Kind because "interface" is a reserved Go keyword
		r.Kind = "NetworkInterface"
		r.References["network_id"] = config.Reference{
			TerraformName: "stackit_network",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_id",true)`,
		}
	})


	p.AddResourceConfigurator("stackit_public_ip", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "stackit_network_interface",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_interface_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_public_ip_associate", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["public_ip_id"] = config.Reference{
			TerraformName: "stackit_public_ip",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("public_ip_id",true)`,
		}
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "stackit_network_interface",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_interface_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_routing_table", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_area_id"] = config.Reference{
			TerraformName: "stackit_network_area",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_area_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_routing_table_route", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["routing_table_id"] = config.Reference{
			TerraformName: "stackit_routing_table",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("routing_table_id",true)`,
		}
		r.References["network_area_id"] = config.Reference{
			TerraformName: "stackit_network_area",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_area_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_security_group", func(r *config.Resource) {
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("stackit_security_group_rule", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "stackit_security_group",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("security_group_id",true)`,
		}
	})
}
