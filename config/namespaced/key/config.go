package key

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the key group (namespaced).
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("litellm_key", func(r *ujconfig.Resource) {
		r.ShortGroup = "key"
		r.Kind = "Key"

		// Mark key and id as secrets, because the LiteLLM provider doesn't do it automatically.
		if s, ok := r.TerraformResource.Schema["id"]; ok {
			s.Sensitive = true
		}
		if s, ok := r.TerraformResource.Schema["key"]; ok {
			s.Sensitive = true
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["key_id"] = []byte(a)
			}
			if a, ok := attr["key"].(string); ok {
				conn["key_value"] = []byte(a)
			}
			return conn, nil
		}
	})
}
