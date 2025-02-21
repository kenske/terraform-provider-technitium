package provider

import (
	"context"
	"fmt"
	"terraform-provider-technitium-dns/internal/technitium"

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
			"Unexpected Data Source Configure Type",
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
		Attributes: map[string]schema.Attribute{
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
		},
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

	var scope technitium.DhcpScope

	// Set values from plan
	scope.Name = plan.Name.ValueString()
	scope.Enabled = plan.Enabled.ValueBool()
	scope.StartingAddress = plan.StartingAddress.ValueString()
	scope.EndingAddress = plan.EndingAddress.ValueString()
	scope.SubnetMask = plan.SubnetMask.ValueString()
	scope.NetworkAddress = plan.NetworkAddress.ValueString()
	scope.BroadcastAddress = plan.BroadcastAddress.ValueString()
	scope.InterfaceAddress = plan.InterfaceAddress.ValueString()

	// Create new order
	_, err := r.client.CreateScope(scope)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dhcp scope",
			"Could not create dhcp scope, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Read refreshes the Terraform state with the latest data.
func (r *dhcpScopeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state dhcpScope
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	scope, err := r.client.GetScope(state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DHCP scope",
			"Could not read DHCP scope  "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite the state with the latest data
	state.Name = types.StringValue(scope.Name)
	state.Enabled = types.BoolValue(scope.Enabled)
	state.StartingAddress = types.StringValue(scope.StartingAddress)
	state.EndingAddress = types.StringValue(scope.EndingAddress)
	state.SubnetMask = types.StringValue(scope.SubnetMask)
	state.NetworkAddress = types.StringValue(scope.NetworkAddress)
	state.BroadcastAddress = types.StringValue(scope.BroadcastAddress)
	state.InterfaceAddress = types.StringValue(scope.InterfaceAddress)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *dhcpScopeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dhcpScopeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
