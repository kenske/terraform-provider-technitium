package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"terraform-provider-technitium/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource               = &dhcpReservedLeaseResource{}
	_ resource.ResourceWithConfigure  = &dhcpReservedLeaseResource{}
	_ resource.ResourceWithModifyPlan = &dhcpReservedLeaseResource{}
)

func NewDhcpReservedLeaseResource() resource.Resource {
	return &dhcpReservedLeaseResource{}
}

type dhcpReservedLeaseResource struct {
	client *technitium.Client
}

// Configure adds the provider configured client to the resource.
func (r *dhcpReservedLeaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*technitium.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected configure type",
			fmt.Sprintf("Expected *technitium.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *dhcpReservedLeaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp_reserved_lease"
}

func (r *dhcpReservedLeaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DhcpReservedLeaseResourceSchema(),
	}

}

func (r *dhcpReservedLeaseResource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	resp.RequiresReplace = path.Paths{
		path.Root("name"),
		path.Root("hardware_address"),
		path.Root("ip_address"),
		path.Root("host_name"),
		path.Root("comments"),
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *dhcpReservedLeaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan dhcpReservedLease
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.SetReservedLease(plan, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating DHCP reserved lease",
			"Could not create DHCP reserved lease: "+err.Error(),
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
func (r *dhcpReservedLeaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *dhcpReservedLeaseResource) SetReservedLease(plan dhcpReservedLease, ctx context.Context) error {

	var lease technitium.DhcpReservedLease

	// Set values from plan
	lease.Name = plan.Name.ValueString()
	lease.HardwareAddress = plan.HardwareAddress.ValueString()
	lease.IpAddress = plan.IpAddress.ValueString()
	lease.HostName = plan.HostName.ValueString()
	lease.Comments = plan.Comments.ValueString()

	err := r.client.CreateLease(lease, ctx)
	if err != nil {
		return err
	}

	return nil
}

// Read refreshes the Terraform state with the latest data.
func (r *dhcpReservedLeaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dhcpReservedLeaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state dhcpReservedLease
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteLease(state.Name.ValueString(), state.HardwareAddress.ValueString(), ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DHCP reserved lease",
			"Could not delete DHCP reserved lease  "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}
}
