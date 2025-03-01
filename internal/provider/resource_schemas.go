package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DhcpScopeResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"starting_address": schema.StringAttribute{
			Required: true,
		},
		"ending_address": schema.StringAttribute{
			Required: true,
		},
		"subnet_mask": schema.StringAttribute{
			Required: true,
		},
		"lease_time_days": schema.Int32Attribute{
			Optional: true,
		},
		"lease_time_hours": schema.Int32Attribute{
			Optional: true,
		},
		"lease_time_minutes": schema.Int32Attribute{
			Optional: true,
		},
		"offer_delay_time": schema.Int32Attribute{
			Optional: true,
		},
		"ping_check_enabled": schema.BoolAttribute{
			Optional: true,
		},
		"ping_check_timeout": schema.Int32Attribute{
			Optional: true,
		},
		"ping_check_retries": schema.Int32Attribute{
			Optional: true,
		},
		"domain_name": schema.StringAttribute{
			Optional: true,
		},
		"domain_search_list": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"dns_updates": schema.BoolAttribute{
			Optional: true,
		},
		"dns_ttl": schema.Int32Attribute{
			Optional: true,
		},
		"server_address": schema.StringAttribute{
			Optional: true,
		},
		"server_host_name": schema.StringAttribute{
			Optional: true,
		},
		"boot_file_name": schema.StringAttribute{
			Optional: true,
		},
		"router_address": schema.StringAttribute{
			Optional: true,
		},
		"use_this_dns_server": schema.BoolAttribute{
			Optional: true,
		},
		"dns_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"wins_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"ntp_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"static_routes": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"destination": schema.StringAttribute{
						Required: true,
					},
					"subnet_mask": schema.StringAttribute{
						Required: true,
					},
					"router": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"vendor_info": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"identifier": schema.StringAttribute{
						Required: true,
					},
					"information": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"capwap_ac_ip_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"tftp_server_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"generic_options": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"code": schema.Int32Attribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"exclusions": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"starting_address": schema.StringAttribute{
						Required: true,
					},
					"ending_address": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"reserved_leases": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"host_name": schema.StringAttribute{
						Optional: true,
					},
					"hardware_address": schema.StringAttribute{
						Required: true,
					},
					"address": schema.StringAttribute{
						Required: true,
					},
					"comments": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"allow_only_reserved_leases": schema.BoolAttribute{
			Optional: true,
		},
		"block_locally_administered_mac_addresses": schema.BoolAttribute{
			Optional: true,
		},
		"ignore_client_identifier_option": schema.BoolAttribute{
			Optional: true,
		},
	}
}
