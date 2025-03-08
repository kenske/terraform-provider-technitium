package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DhcpScopesSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Computed: true,
		},
		"enabled": schema.BoolAttribute{
			Computed: true,
		},
		"starting_address": schema.StringAttribute{
			Computed: true,
		},
		"ending_address": schema.StringAttribute{
			Computed: true,
		},
		"subnet_mask": schema.StringAttribute{
			Computed: true,
		},
		"network_address": schema.StringAttribute{
			Computed: true,
		},
		"broadcast_address": schema.StringAttribute{
			Computed: true,
		},
	}
}

func DhcpScopeSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"starting_address": schema.StringAttribute{
			Computed: true,
		},
		"ending_address": schema.StringAttribute{
			Computed: true,
		},
		"subnet_mask": schema.StringAttribute{
			Computed: true,
		},
		"router_address": schema.StringAttribute{
			Computed: true,
		},
	}
}
