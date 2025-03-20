package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type dhcpScopeList struct {
	Name             types.String `tfsdk:"name"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	StartingAddress  types.String `tfsdk:"starting_address"`
	EndingAddress    types.String `tfsdk:"ending_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	NetworkAddress   types.String `tfsdk:"network_address"`
	BroadcastAddress types.String `tfsdk:"broadcast_address"`
}

type dhcpScope struct {
	Name             types.String   `tfsdk:"name"`
	StartingAddress  types.String   `tfsdk:"starting_address"`
	EndingAddress    types.String   `tfsdk:"ending_address"`
	SubnetMask       types.String   `tfsdk:"subnet_mask"`
	RouterAddress    types.String   `tfsdk:"router_address"`
	UseThisDnsServer types.Bool     `tfsdk:"use_this_dns_server"`
	DnsServers       []types.String `tfsdk:"dns_servers"`
	DomainName       types.String   `tfsdk:"domain_name"`
	Exclusions       []Exclusion    `tfsdk:"exclusions"`
}

type Exclusion struct {
	StartingAddress types.String `tfsdk:"starting_address"`
	EndingAddress   types.String `tfsdk:"ending_address"`
}

type dhcpReservedLease struct {
	Name            types.String `tfsdk:"name"`
	HardwareAddress types.String `tfsdk:"hardware_address"`
	IpAddress       types.String `tfsdk:"ip_address"`
	HostName        types.String `tfsdk:"host_name"`
	Comments        types.String `tfsdk:"comments"`
}
