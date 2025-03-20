package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
		"use_this_dns_server": schema.BoolAttribute{
			Computed: true,
		},
		"dns_servers": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"domain_name": schema.StringAttribute{
			Computed: true,
		},
		"exclusions": schema.ListAttribute{
			Computed: true,
			ElementType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"starting_address": types.StringType,
					"ending_address":   types.StringType,
				},
			},
		},
	}
}
