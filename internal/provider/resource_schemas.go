package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DhcpScopeResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required:    true,
			Description: "The name of the DHCP scope.",
		},
		"starting_address": schema.StringAttribute{
			Required:    true,
			Description: "The starting IP address of the DHCP scope.",
		},
		"ending_address": schema.StringAttribute{
			Required:    true,
			Description: "The ending IP address of the DHCP scope.",
		},
		"subnet_mask": schema.StringAttribute{
			Required:    true,
			Description: "The subnet mask of the network.",
		},
		"lease_time_days": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(1),
			Description: "The lease time in days. Default is 1 day.",
		},
		"lease_time_hours": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(0),
			Description: "The lease time in hours.",
		},
		"lease_time_minutes": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(0),
			Description: "The lease time in minutes.",
		},
		"offer_delay_time": schema.Int64Attribute{
			Optional:    true,
			Description: "The time duration in milliseconds that the DHCP server delays sending an offer.",
		},
		"ping_check_enabled": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Enable ping check to ensure IP address is available before offering it.",
		},
		"ping_check_timeout": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(1000),
			Description: "The timeout interval in milliseconds to wait for ping check response.",
		},
		"ping_check_retries": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(2),
			Description: "The number of retries for ping check.",
		},
		"router_address": schema.StringAttribute{
			Optional:    true,
			Description: "The default gateway IP address.",
		},
		"use_this_dns_server": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Use this DNS server's IP address for the DNS server option.",
		},
		"dns_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of DNS server IP addresses to use.",
		},
		"wins_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of WINS server IP addresses.",
		},
		"ntp_servers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of NTP server IP addresses.",
		},
		"ntp_server_domain_names": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of NTP server domain names.",
		},
		"static_routes": schema.ListNestedAttribute{
			Optional:    true,
			Description: "List of classless static routes.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"destination": schema.StringAttribute{
						Required:    true,
						Description: "The destination network address.",
					},
					"subnet_mask": schema.StringAttribute{
						Required:    true,
						Description: "The subnet mask for the destination.",
					},
					"router": schema.StringAttribute{
						Required:    true,
						Description: "The router IP address for this route.",
					},
				},
			},
		},
		"vendor_info": schema.ListNestedAttribute{
			Optional:    true,
			Description: "List of vendor specific information options.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"identifier": schema.StringAttribute{
						Required:    true,
						Description: "The vendor class identifier string.",
					},
					"information": schema.StringAttribute{
						Required:    true,
						Description: "The vendor specific information as a hex string.",
					},
				},
			},
		},
		"capwap_ac_ip_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of CAPWAP Access Controller IP addresses for wireless access points.",
		},
		"tftp_server_addresses": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of TFTP server IP addresses.",
		},
		"generic_options": schema.ListNestedAttribute{
			Optional:    true,
			Description: "List of generic DHCP options.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"code": schema.Int64Attribute{
						Required:    true,
						Description: "The DHCP option code.",
					},
					"value": schema.StringAttribute{
						Required:    true,
						Description: "The DHCP option value as a hex string.",
					},
				},
			},
		},
		"exclusions": schema.ListNestedAttribute{
			Optional:    true,
			Description: "List of IP address ranges to exclude from the scope.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"starting_address": schema.StringAttribute{
						Required:    true,
						Description: "The starting IP address of the exclusion range.",
					},
					"ending_address": schema.StringAttribute{
						Required:    true,
						Description: "The ending IP address of the exclusion range.",
					},
				},
			},
		},
		"reserved_leases": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of reserved lease scope names.",
		},
		"allow_only_reserved_leases": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Allow only reserved leases from this scope.",
		},
		"block_locally_administered_mac_addresses": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Block locally administered MAC addresses from getting leases.",
		},
		"ignore_client_identifier_option": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Ignore client identifier option and use hardware address instead.",
		},
		"domain_name": schema.StringAttribute{
			Optional:    true,
			Description: "The domain name to be used by the client.",
		},
		"domain_search_list": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "List of domain names for DNS resolution.",
		},
		"boot_file_name": schema.StringAttribute{
			Optional:    true,
			Description: "The boot file name for network booting.",
		},
		"next_server_address": schema.StringAttribute{
			Optional:    true,
			Description: "The next server IP address for network booting.",
		},
		"server_host_name": schema.StringAttribute{
			Optional:    true,
			Description: "The DHCP server host name.",
		},
		"server_address": schema.StringAttribute{
			Optional:    true,
			Description: "The DHCP server IP address.",
		},
		"interface_address": schema.StringAttribute{
			Optional:    true,
			Description: "The network interface IP address to bind the scope.",
		},
		"interface_index": schema.Int64Attribute{
			Optional:    true,
			Description: "The network interface index to bind the scope.",
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
				"[`Primary`, `Secondary`, `Stub`, `Forwarder`, `SecondaryForwarder`, `Catalog`, `SecondaryCatalog`]",
		},
		"catalog": schema.StringAttribute{
			Optional: true,
		},
		"forwarder": schema.StringAttribute{
			Optional: true,
			Description: "The address of the DNS server to be used as a forwarder. " +
				"This optional parameter is required to be used with Conditional Forwarder zones. " +
				"A special value `this-server` can be used as a forwarder which when used will " +
				"forward all the requests internally to this DNS server such that you can override " +
				"the zone with records and rest of the zone gets resolved via this server. " +
				"The `initialize_forwarder` parameter must be set to `true` to use this option.",
		},
		"initialize_forwarder": schema.BoolAttribute{
			Optional: true,
			Computed: true,
			Description: " Set value as true to initialize the Conditional Forwarder zone with " +
				"an FWD record or set it to false to create an empty Forwarder zone. " +
				"Default value is `true`",
			Default: booldefault.StaticBool(true),
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
			Description: "The DNS transport protocol to be used by the Conditional Forwarder zone." +
				"This optional parameter is used with Conditional Forwarder zones." +
				"Valid values are [`Udp`, `Tcp`, `Tls`, `Https`, `Quic`]." +
				"Default is `Udp`",
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
			Required:    true,
			Description: "The type of DNS record. Valid values are [`A`, `AAAA`, `NS`, `CNAME`, `PTR`, `MX`, `TXT`, `SRV`, `DNAME`, `DS`, `SSHFP`, `TLSA`, `SVCB`, `HTTPS`, `URI`, `CAA`] and proprietary types [`ANAME`, `FWD`, `APP`].",
		},
		"zone": schema.StringAttribute{
			Optional: true,
		},
		"ttl": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The time-to-live (TTL) for the DNS record in seconds",
		},
		"disabled": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			Description: "Set to true to disable the DNS record. Default is false.",
		},
		"comments": schema.StringAttribute{
			Optional: true,
			Computed: true,
			Default:  stringdefault.StaticString(""),
		},
		"expiry_ttl": schema.Int64Attribute{
			Optional: true,
			Computed: true,
			Default:  int64default.StaticInt64(0),
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
			Description: "This parameter is required for adding the FWD record. Valid values are [`Udp`, `Tcp`, `Tls`, `Https`, `Quic`].",
		},
		"forwarder": schema.StringAttribute{
			Optional:    true,
			Description: "The forwarder address. A special value of `this-server` can be used to directly forward requests internally to the DNS server. This parameter is required for adding the FWD record.",
		},
		"forwarder_priority": schema.Int64Attribute{
			Optional: true,
		},
		"dnssec_validation": schema.BoolAttribute{
			Optional: true,
		},
		"proxy_type": schema.StringAttribute{
			Optional:    true,
			Description: "The type of proxy to be used for conditional forwarding. Valid values are [`NoProxy`, `DefaultProxy`, `Http`, `Socks5`].",
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
		"app_name": schema.StringAttribute{
			Optional:    true,
			Description: "DNS app name, required for `APP` records",
		},
		"class_path": schema.StringAttribute{
			Optional:    true,
			Description: "DNS app class path",
		},
		"record_data": schema.StringAttribute{
			Optional:    true,
			Description: "DNS app record data",
		},
	}
}
