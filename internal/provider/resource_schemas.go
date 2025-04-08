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
		"domain_name": schema.StringAttribute{
			Optional: true,
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
	}
}

func DhcpReservedLeaseResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{

		"name": schema.StringAttribute{
			Required: true,
		},
		"hardware_address": schema.StringAttribute{
			Required: true,
		},
		"ip_address": schema.StringAttribute{
			Required: true,
		},
		"host_name": schema.StringAttribute{
			Optional: true,
		},
		"comments": schema.StringAttribute{
			Optional: true,
		},
	}
}

func DnsZoneResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Required: true,
		},
		"catalog": schema.StringAttribute{
			Optional: true,
		},
		"forwarder": schema.StringAttribute{
			Optional:  true,
			WriteOnly: true,
		},
		"use_soa_serial_date_scheme": schema.BoolAttribute{
			Optional:  true,
			WriteOnly: true,
		},
		"primary_name_server_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			WriteOnly:   true,
		},
		"zone_transfer_protocol": schema.StringAttribute{
			Optional:  true,
			WriteOnly: true,
		},
		"tsig_key_name": schema.StringAttribute{
			Optional:  true,
			WriteOnly: true,
		},
		"protocol": schema.StringAttribute{
			Optional:  true,
			WriteOnly: true,
		},
		"dnssec_validation": schema.BoolAttribute{
			Optional:  true,
			WriteOnly: true,
		},
	}
}
