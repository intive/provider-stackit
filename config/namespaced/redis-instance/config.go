package redisinstance

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_redis_instance", func(r *config.Resource) {
		r.ShortGroup = "redis"
	})
}
