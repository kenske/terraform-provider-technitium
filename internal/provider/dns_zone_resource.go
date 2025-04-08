package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"terraform-provider-technitium-dns/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource               = &dnsZoneResource{}
	_ resource.ResourceWithConfigure  = &dnsZoneResource{}
	_ resource.ResourceWithModifyPlan = &dnsZoneResource{}
)

func NewDnsZoneResource() resource.Resource {
	return &dnsZoneResource{}
}

type dnsZoneResource struct {
	client *technitium.Client
}

// Configure adds the provider configured client to the resource.
func (r *dnsZoneResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = ConfigureClient(req, resp)
}

// Metadata returns the resource type name.
func (r *dnsZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone"
}

func (r *dnsZoneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DnsZoneResourceSchema(),
	}

}

// Create creates the resource and sets the initial Terraform state.
func (r *dnsZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan dnsZoneCreate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.CreateZone(plan, ctx)
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
func (r *dnsZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *dnsZoneResource) CreateZone(plan dnsZoneCreate, ctx context.Context) error {

	var zone technitium.DnsZoneCreate

	// Set values from plan
	zone.Name = plan.Name.ValueString()
	zone.Type = plan.Type.ValueString()
	zone.Catalog = plan.Catalog.ValueString()
	zone.Forwarder = plan.Forwarder.ValueString()
	zone.UseSoaSerialDateScheme = plan.UseSoaSerialDateScheme.ValueBool()
	zone.PrimaryNameServerAddresses = convertTfListToStringList(plan.PrimaryNameServerAddresses)
	zone.ZoneTransferProtocol = plan.ZoneTransferProtocol.ValueString()
	zone.TsigKeyName = plan.TsigKeyName.ValueString()
	zone.Protocol = plan.Protocol.ValueString()
	zone.DnssecValidation = plan.DnssecValidation.ValueBool()

	// Create new scope
	_, err := r.client.CreateDnsZone(zone, ctx)
	if err != nil {
		return err
	}

	return nil
}

// Read refreshes the Terraform state with the latest data.
func (r *dnsZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state dnsZoneCreate
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, err := r.client.GetDnsZone(state.Name.ValueString(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DNS zone scope",
			"Could not read DNS zone  "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite the state with the latest data
	state.Name = types.StringValue(zone.Name)
	state.Type = types.StringValue(zone.Type)
	state.Catalog = types.StringValue(zone.Catalog)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dnsZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

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

func (r *dnsZoneResource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	resp.RequiresReplace = path.Paths{
		path.Root("name"),
		path.Root("type"),
		path.Root("catalog"),
	}
}
