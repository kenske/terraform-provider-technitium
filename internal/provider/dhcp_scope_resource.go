package provider

import (
	"context"
	"fmt"
	"terraform-provider-technitium/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &dhcpScopeResource{}
	_ resource.ResourceWithConfigure = &dhcpScopeResource{}
)

func NewDhcpScopeResource() resource.Resource {
	return &dhcpScopeResource{}
}

type dhcpScopeResource struct {
	client *technitium.Client
}

// Configure adds the provider configured client to the resource.
func (r *dhcpScopeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*technitium.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Configure Type",
			fmt.Sprintf("Expected *technitium.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *dhcpScopeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp_scope"
}

func (r *dhcpScopeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes:  DhcpScopeResourceSchema(),
		Description: "The API does a lot of magic behind the scenes when creating a new scope, so most changes will require deleting and recreating the scope.",
	}

}

// Create creates the resource and sets the initial Terraform state.
func (r *dhcpScopeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan dhcpScope
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.SetScope(plan, "", ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating DHCP scope",
			"Could not create DHCP scope: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Update updates the resource and sets the updated Terraform state on success.
func (r *dhcpScopeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan dhcpScope
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get state so that we can pass the old name to the SetScope method
	var state dhcpScope
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	oldName := state.Name.ValueString()

	err := r.SetScope(plan, oldName, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating DHCP scope",
			"Could not updating DHCP scope: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *dhcpScopeResource) SetScope(plan dhcpScope, oldName string, ctx context.Context) error {

	var scope technitium.DhcpScope

	// Basic scope configuration
	scope.Name = plan.Name.ValueString()
	scope.StartingAddress = plan.StartingAddress.ValueString()
	scope.EndingAddress = plan.EndingAddress.ValueString()
	scope.SubnetMask = plan.SubnetMask.ValueString()

	// Lease time configuration
	if !plan.LeaseTimeDays.IsNull() {
		scope.LeaseTimeDays = int(plan.LeaseTimeDays.ValueInt64())
	}
	if !plan.LeaseTimeHours.IsNull() {
		scope.LeaseTimeHours = int(plan.LeaseTimeHours.ValueInt64())
	}
	if !plan.LeaseTimeMinutes.IsNull() {
		scope.LeaseTimeMinutes = int(plan.LeaseTimeMinutes.ValueInt64())
	}

	// Offer delay and ping check
	if !plan.OfferDelayTime.IsNull() {
		scope.OfferDelayTime = int(plan.OfferDelayTime.ValueInt64())
	}
	if !plan.PingCheckEnabled.IsNull() {
		scope.PingCheckEnabled = plan.PingCheckEnabled.ValueBool()
	}
	if !plan.PingCheckTimeout.IsNull() {
		scope.PingCheckTimeout = int(plan.PingCheckTimeout.ValueInt64())
	}
	if !plan.PingCheckRetries.IsNull() {
		scope.PingCheckRetries = int(plan.PingCheckRetries.ValueInt64())
	}

	// Network configuration
	if !plan.RouterAddress.IsNull() {
		scope.RouterAddress = plan.RouterAddress.ValueString()
	}
	if !plan.UseThisDnsServer.IsNull() {
		scope.UseThisDnsServer = plan.UseThisDnsServer.ValueBool()
	}

	// DNS servers
	if len(plan.DnsServers) > 0 {
		scope.DnsServers = make([]string, 0, len(plan.DnsServers))
		for _, server := range plan.DnsServers {
			scope.DnsServers = append(scope.DnsServers, server.ValueString())
		}
	}

	// WINS servers
	if len(plan.WinsServers) > 0 {
		scope.WinsServers = make([]string, 0, len(plan.WinsServers))
		for _, server := range plan.WinsServers {
			scope.WinsServers = append(scope.WinsServers, server.ValueString())
		}
	}

	// NTP servers
	if len(plan.NtpServers) > 0 {
		scope.NtpServers = make([]string, 0, len(plan.NtpServers))
		for _, server := range plan.NtpServers {
			scope.NtpServers = append(scope.NtpServers, server.ValueString())
		}
	}

	// NTP server domain names
	if len(plan.NtpServerDomainNames) > 0 {
		scope.NtpServerDomainNames = make([]string, 0, len(plan.NtpServerDomainNames))
		for _, name := range plan.NtpServerDomainNames {
			scope.NtpServerDomainNames = append(scope.NtpServerDomainNames, name.ValueString())
		}
	}

	// Static routes
	if len(plan.StaticRoutes) > 0 {
		scope.StaticRoutes = make([]technitium.DhcpStaticRoute, 0, len(plan.StaticRoutes))
		for _, route := range plan.StaticRoutes {
			scope.StaticRoutes = append(scope.StaticRoutes, technitium.DhcpStaticRoute{
				Destination: route.Destination.ValueString(),
				SubnetMask:  route.SubnetMask.ValueString(),
				Router:      route.Router.ValueString(),
			})
		}
	}

	// Vendor info
	if len(plan.VendorInfo) > 0 {
		scope.VendorInfo = make([]technitium.DhcpVendorInfo, 0, len(plan.VendorInfo))
		for _, vendor := range plan.VendorInfo {
			scope.VendorInfo = append(scope.VendorInfo, technitium.DhcpVendorInfo{
				Identifier:  vendor.Identifier.ValueString(),
				Information: vendor.Information.ValueString(),
			})
		}
	}

	// CAPWAP AC IP addresses
	if len(plan.CAPWAPAcIpAddresses) > 0 {
		scope.CAPWAPAcIpAddresses = make([]string, 0, len(plan.CAPWAPAcIpAddresses))
		for _, addr := range plan.CAPWAPAcIpAddresses {
			scope.CAPWAPAcIpAddresses = append(scope.CAPWAPAcIpAddresses, addr.ValueString())
		}
	}

	// TFTP server addresses
	if len(plan.TftpServerAddresses) > 0 {
		scope.TftpServerAddresses = make([]string, 0, len(plan.TftpServerAddresses))
		for _, addr := range plan.TftpServerAddresses {
			scope.TftpServerAddresses = append(scope.TftpServerAddresses, addr.ValueString())
		}
	}

	// Generic options
	if len(plan.GenericOptions) > 0 {
		scope.GenericOptions = make([]technitium.DhcpGenericOption, 0, len(plan.GenericOptions))
		for _, option := range plan.GenericOptions {
			scope.GenericOptions = append(scope.GenericOptions, technitium.DhcpGenericOption{
				Code:  int(option.Code.ValueInt64()),
				Value: option.Value.ValueString(),
			})
		}
	}

	// Exclusions
	if len(plan.Exclusions) > 0 {
		scope.Exclusions = make([]technitium.Exclusion, 0, len(plan.Exclusions))
		for _, e := range plan.Exclusions {
			scope.Exclusions = append(scope.Exclusions, technitium.Exclusion{
				StartingAddress: e.StartingAddress.ValueString(),
				EndingAddress:   e.EndingAddress.ValueString(),
			})
		}
	}

	// Reserved leases - convert from string list to DhcpReservedLease objects
	if len(plan.ReservedLeases) > 0 {
		scope.ReservedLeases = make([]technitium.DhcpReservedLease, 0, len(plan.ReservedLeases))
		for _, lease := range plan.ReservedLeases {
			scope.ReservedLeases = append(scope.ReservedLeases, technitium.DhcpReservedLease{
				HardwareAddress: lease.ValueString(),
			})
		}
	}

	// Access control options
	if !plan.AllowOnlyReservedLeases.IsNull() {
		scope.AllowOnlyReservedLeases = plan.AllowOnlyReservedLeases.ValueBool()
	}
	if !plan.BlockLocallyAdministeredMacAddresses.IsNull() {
		scope.BlockLocallyAdministeredMacAddresses = plan.BlockLocallyAdministeredMacAddresses.ValueBool()
	}
	if !plan.IgnoreClientIdentifierOption.IsNull() {
		scope.IgnoreClientIdentifierOption = plan.IgnoreClientIdentifierOption.ValueBool()
	}

	// Domain configuration
	if !plan.DomainName.IsNull() {
		scope.DomainName = plan.DomainName.ValueString()
	}
	if len(plan.DomainSearchList) > 0 {
		scope.DomainSearchList = make([]string, 0, len(plan.DomainSearchList))
		for _, domain := range plan.DomainSearchList {
			scope.DomainSearchList = append(scope.DomainSearchList, domain.ValueString())
		}
	}

	// Boot options
	if !plan.BootFileName.IsNull() {
		scope.BootFileName = plan.BootFileName.ValueString()
	}
	if !plan.NextServerAddress.IsNull() {
		scope.NextServerAddress = plan.NextServerAddress.ValueString()
	}
	if !plan.ServerHostName.IsNull() {
		scope.ServerHostName = plan.ServerHostName.ValueString()
	}
	if !plan.ServerAddress.IsNull() {
		scope.ServerAddress = plan.ServerAddress.ValueString()
	}

	// Interface binding
	if !plan.InterfaceAddress.IsNull() {
		scope.InterfaceAddress = plan.InterfaceAddress.ValueString()
	}
	if !plan.InterfaceIndex.IsNull() {
		scope.InterfaceIndex = int(plan.InterfaceIndex.ValueInt64())
	}

	// Create or update scope
	createdScope, err := r.client.SetScope(scope, oldName, ctx)
	if err != nil {
		return err
	}

	// Set state to fully populated data
	plan.Name = types.StringValue(createdScope.Name)
	plan.SubnetMask = types.StringValue(createdScope.SubnetMask)
	plan.StartingAddress = types.StringValue(createdScope.StartingAddress)
	plan.EndingAddress = types.StringValue(createdScope.EndingAddress)

	// Set router address if not empty
	plan.RouterAddress = types.StringNull()
	if createdScope.RouterAddress != "" {
		plan.RouterAddress = types.StringValue(createdScope.RouterAddress)
	}

	return nil
}

// Read refreshes the Terraform state with the latest data.
func (r *dhcpScopeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state dhcpScope
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed scope value from Technitium
	scope, err := r.client.GetScope(state.Name.ValueString(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DHCP scope",
			"Could not read DHCP scope  "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite the state with the latest data
	state.Name = types.StringValue(scope.Name)
	state.StartingAddress = types.StringValue(scope.StartingAddress)
	state.EndingAddress = types.StringValue(scope.EndingAddress)
	state.SubnetMask = types.StringValue(scope.SubnetMask)
	state.RouterAddress = types.StringNull()

	//Set router address if not empty
	if scope.RouterAddress != "" {
		state.RouterAddress = types.StringValue(scope.RouterAddress)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dhcpScopeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state dhcpScope
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteScope(state.Name.ValueString(), ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DHCP scope",
			"Could not delete DHCP scope  "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}
}
