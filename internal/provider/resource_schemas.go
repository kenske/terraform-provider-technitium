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
			Description: "The type of DNS zone. Valid values are " +
				"[Primary, Secondary, Stub, Forwarder, SecondaryForwarder, Catalog, SecondaryCatalog]",
		},
		"catalog": schema.StringAttribute{
			Optional: true,
		},
		"forwarder": schema.StringAttribute{
			Optional: true,
			Description: "The address of the DNS server to be used as a forwarder. " +
				"This optional parameter is required to be used with Conditional Forwarder zones. " +
				"A special value this-server can be used as a forwarder which when used will " +
				"forward all the requests internally to this DNS server such that you can override " +
				"the zone with records and rest of the zone gets resolved via This Server. " +
				"The initialize_forwarder parameter must be set to true to use this option.",
		},
		"initialize_forwarder": schema.BoolAttribute{
			Optional: true,
			Description: " Set value as true to initialize the Conditional Forwarder zone with " +
				"an FWD record or set it to false to create an empty Forwarder zone. " +
				"Default value is true",
		},
		"use_soa_serial_date_scheme": schema.BoolAttribute{
			Optional: true,
		},
		"primary_name_server_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"zone_transfer_protocol": schema.StringAttribute{
			Optional: true,
		},
		"tsig_key_name": schema.StringAttribute{
			Optional: true,
		},
		"protocol": schema.StringAttribute{
			Optional: true,
		},
		"dnssec_validation": schema.BoolAttribute{
			Optional: true,
		},
	}
}

func DnsZoneRecordResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"domain": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Required: true,
		},
		"zone": schema.StringAttribute{
			Optional: true,
		},
		"ttl": schema.Int64Attribute{
			Optional:    true,
			Description: "The time-to-live (TTL) for the DNS record in seconds",
		},
		"comments": schema.StringAttribute{
			Optional: true,
		},
		"expiry_ttl": schema.Int64Attribute{
			Optional: true,
		},
		"ip_address": schema.StringAttribute{
			Optional: true,
		},
		"ptr": schema.StringAttribute{
			Optional: true,
		},
		"create_ptr_zone": schema.BoolAttribute{
			Optional: true,
		},
		"update_svcb_hints": schema.BoolAttribute{
			Optional: true,
		},
		"name_server": schema.StringAttribute{
			Optional: true,
		},
		"cname": schema.StringAttribute{
			Optional: true,
		},
		"ptr_name": schema.StringAttribute{
			Optional: true,
		},
		"exchange": schema.StringAttribute{
			Optional:    true,
			Description: "The exchange domain name. This option is required for adding MX record.",
		},
		"preference": schema.Int64Attribute{
			Optional:    true,
			Description: "This is the preference value for MX record type. This option is required for adding MX record.",
		},
		"text": schema.StringAttribute{
			Optional:    true,
			Description: "The text data for TXT record. This option is required for adding TXT record.",
		},
		"split_text": schema.StringAttribute{
			Optional:    true,
			Description: "Set to true for using new line char to split text into multiple character-strings for adding TXT record.",
		},
		"protocol": schema.StringAttribute{
			Optional:    true,
			Description: "This parameter is required for adding the FWD record. Valid values are [Udp, Tcp, Tls, Https, Quic].",
		},
		"forwarder": schema.StringAttribute{
			Optional:    true,
			Description: "The forwarder address. A special value of this-server can be used to directly forward requests internally to the DNS server. This parameter is required for adding the FWD record.",
		},
		"forwarder_priority": schema.Int64Attribute{
			Optional: true,
		},
		"dnssec_validation": schema.BoolAttribute{
			Optional: true,
		},
		"proxy_type": schema.StringAttribute{
			Optional: true,
		},
		"proxy_address": schema.StringAttribute{
			Optional: true,
		},
		"proxy_port": schema.Int64Attribute{
			Optional: true,
		},
		"proxy_username": schema.StringAttribute{
			Optional: true,
		},
		"proxy_password": schema.StringAttribute{
			Optional: true,
		},
	}
}
