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
		Attributes: DhcpScopeResourceSchema(),
	}

}

// Create creates the resource and sets the initial Terraform state.
func (r *dhcpScopeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan dhcpScopeCreate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var scope technitium.DhcpScope

	// Set values from plan
	scope.Name = plan.Name.ValueString()
	scope.StartingAddress = plan.StartingAddress.ValueString()
	scope.EndingAddress = plan.EndingAddress.ValueString()
	scope.SubnetMask = plan.SubnetMask.ValueString()

	// TODO: Ternary operation to keep null value instead of empty string
	scope.RouterAddress = plan.RouterAddress.IsNull() ? nil : plan.RouterAddress.ValueString()

	// Create new order
	_, err := r.client.CreateScope(scope, ctx)
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

	var state dhcpScopeGet
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
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
	state.RouterAddress = types.StringValue(scope.RouterAddress)

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

	var state dhcpScopeGet
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
