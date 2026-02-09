package key

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the key group (namespaced).
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("litellm_key", func(r *ujconfig.Resource) {
		r.ShortGroup = "key"
		r.Kind = "Key"
	})
}
