package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type dhcpScopeGet struct {
	Name            types.String `tfsdk:"name"`
	StartingAddress types.String `tfsdk:"starting_address"`
	EndingAddress   types.String `tfsdk:"ending_address"`
	SubnetMask      types.String `tfsdk:"subnet_mask"`
	RouterAddress   types.String `tfsdk:"router_address"`
}

type dhcpScopeList struct {
	Name             types.String `tfsdk:"name"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	StartingAddress  types.String `tfsdk:"starting_address"`
	EndingAddress    types.String `tfsdk:"ending_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	NetworkAddress   types.String `tfsdk:"network_address"`
	BroadcastAddress types.String `tfsdk:"broadcast_address"`
}

type dhcpScopeSet struct {
	Name            types.String `tfsdk:"name"`
	StartingAddress types.String `tfsdk:"starting_address"`
	EndingAddress   types.String `tfsdk:"ending_address"`
	SubnetMask      types.String `tfsdk:"subnet_mask"`
	RouterAddress   types.String `tfsdk:"router_address"`
}
