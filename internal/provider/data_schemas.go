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

func DnsZoneRecordsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"domain": schema.StringAttribute{
			Required: true,
		},
		"records": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: DnsZoneRecordSchema(),
			},
		},
	}
}

func DnsZoneRecordDataSourceSchema() map[string]schema.Attribute {
	// Start with the base schema from DnsZoneRecordSchema.
	schemaAttributes := DnsZoneRecordSchema()

	// For the data source, 'name' is a required input to find the record.
	schemaAttributes["name"] = schema.StringAttribute{
		Required: true,
	}

	// 'type' is also a required input for the lookup.
	schemaAttributes["type"] = schema.StringAttribute{
		Required: true,
	}

	return schemaAttributes
}

func DnsZoneRecordSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Computed: true,
		},
		"type": schema.StringAttribute{
			Computed: true,
		},
		"ttl": schema.Int64Attribute{
			Computed: true,
		},
		"disabled": schema.BoolAttribute{
			Computed: true,
		},
		"dnssec_status": schema.StringAttribute{
			Computed: true,
		},
		"last_used_on": schema.StringAttribute{
			Computed: true,
		},
		"last_modified": schema.StringAttribute{
			Computed: true,
		},
		"expiry_ttl": schema.Int64Attribute{
			Computed: true,
		},
		"record_data": schema.SingleNestedAttribute{
			Computed:   true,
			Attributes: RecordDataSchema(),
		},
	}
}

func RecordDataSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"primary_name_server": schema.StringAttribute{
			Computed: true,
		},
		"responsible_person": schema.StringAttribute{
			Computed: true,
		},
		"serial": schema.Int32Attribute{
			Computed: true,
		},
		"refresh": schema.Int32Attribute{
			Computed: true,
		},
		"retry": schema.Int32Attribute{
			Computed: true,
		},
		"expire": schema.Int32Attribute{
			Computed: true,
		},
		"minimum": schema.Int32Attribute{
			Computed: true,
		},
		"use_serial_date_scheme": schema.BoolAttribute{
			Computed: true,
		},
		"protocol": schema.StringAttribute{
			Computed: true,
		},
		"forwarder": schema.StringAttribute{
			Computed: true,
		},
		"priority": schema.Int32Attribute{
			Computed: true,
		},
		"dnssec_validation": schema.BoolAttribute{
			Computed: true,
		},
		"proxy_type": schema.StringAttribute{
			Computed: true,
		},
		"ip_address": schema.StringAttribute{
			Computed: true,
		},
		"cname": schema.StringAttribute{
			Computed: true,
		},
		"name_server": schema.StringAttribute{
			Computed: true,
		},
	}
}
