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

func DnsZoneSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Computed: true,
		},
		"disabled": schema.BoolAttribute{
			Computed: true,
		},
		"dnssec_status": schema.StringAttribute{
			Computed: true,
		},
		"catalog": schema.StringAttribute{
			Computed: true,
		},
		"notify_failed": schema.BoolAttribute{
			Computed: true,
		},
		"notify_failed_for": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"query_access": schema.StringAttribute{
			Computed: true,
		},
		"query_access_network_acl": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"zone_transfer": schema.StringAttribute{
			Computed: true,
		},
		"zone_transfer_network_acl": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"zone_transfer_tsig_key_names": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"notify": schema.StringAttribute{
			Computed: true,
		},
		"notify_name_servers": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
		"update": schema.StringAttribute{
			Computed: true,
		},
		"update_network_acl": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
		},
	}
}

func DnsZonesSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Computed: true,
		},
		"type": schema.StringAttribute{
			Computed: true,
		},
		"disabled": schema.BoolAttribute{
			Computed: true,
		},
		"dnssec_status": schema.StringAttribute{
			Computed: true,
		},
		"soa_serial": schema.Int32Attribute{
			Computed: true,
		},
		"expiry": schema.StringAttribute{
			Computed: true,
		},
		"is_expired": schema.BoolAttribute{
			Computed: true,
		},
		"last_modified": schema.StringAttribute{
			Computed: true,
		},
		"internal": schema.BoolAttribute{
			Computed: true,
		},
		"catalog": schema.StringAttribute{
			Computed: true,
		},
	}
}
