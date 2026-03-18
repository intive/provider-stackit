package compute

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_affinity_group", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("stackit_image", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("stackit_key_pair", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("stackit_volume", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("stackit_server", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["image_id"] = config.Reference{
			TerraformName: "stackit_image",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("image_id",true)`,
		}
		r.References["network_interfaces"] = config.Reference{
			TerraformName: "stackit_network_interface",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_interface_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_server_network_interface_attach", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "stackit_server",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("server_id",true)`,
		}
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "stackit_network_interface",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("network_interface_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_server_backup_schedule", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "stackit_server",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("server_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_server_service_account_attach", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "stackit_server",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("server_id",true)`,
		}
		r.References["service_account_email"] = config.Reference{
			TerraformName: "stackit_service_account",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("email",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_server_update_schedule", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "stackit_server",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("server_id",true)`,
		}
	})

	p.AddResourceConfigurator("stackit_server_volume_attach", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "stackit_server",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("server_id",true)`,
		}
		r.References["volume_id"] = config.Reference{
			TerraformName: "stackit_volume",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("volume_id",true)`,
		}
	})
}
