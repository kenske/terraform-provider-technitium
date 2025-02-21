package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type dhcpScope struct {
	Name             types.String `tfsdk:"name"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	StartingAddress  types.String `tfsdk:"starting_address"`
	EndingAddress    types.String `tfsdk:"ending_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	NetworkAddress   types.String `tfsdk:"network_address"`
	BroadcastAddress types.String `tfsdk:"broadcast_address"`
	InterfaceAddress types.String `tfsdk:"interface_address"`
}

type dhcpScopeResourceModel struct {
	dhcpScope
}
