package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-technitium/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource               = &dnsZoneRecordResource{}
	_ resource.ResourceWithConfigure  = &dnsZoneRecordResource{}
	_ resource.ResourceWithModifyPlan = &dnsZoneRecordResource{}
)

func NewDnsZoneRecordResource() resource.Resource {
	return &dnsZoneRecordResource{}
}

type dnsZoneRecordResource struct {
	client *technitium.Client
}

// Configure adds the provider configured client to the resource.
func (r *dnsZoneRecordResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = ConfigureResourceClient(req, resp)
}

// Metadata returns the resource type name.
func (r *dnsZoneRecordResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone_record"
}

func (r *dnsZoneRecordResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DnsZoneRecordResourceSchema(),
	}

}

// Create creates the resource and sets the initial Terraform state.
func (r *dnsZoneRecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan dnsZoneRecordCreate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.CreateZoneRecord(plan, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating zone record",
			"Could not create zone record: "+err.Error(),
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
func (r *dnsZoneRecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *dnsZoneRecordResource) CreateZoneRecord(plan dnsZoneRecordCreate, ctx context.Context) error {

	var record technitium.DnsZoneRecordCreate

	record.Domain = plan.Domain.ValueString()
	record.Type = plan.Type.ValueString()
	record.Zone = plan.Zone.ValueString()
	record.TTL = plan.TTL.ValueInt64()
	record.Comments = plan.Comments.ValueString()
	record.ExpiryTTL = plan.ExpiryTTL.ValueInt64()
	record.IPAddress = plan.IPAddress.ValueString()
	record.Ptr = plan.Ptr.ValueString()
	record.CreatePtrZone = plan.CreatePtrZone.ValueBool()
	record.UpdateSvcbHints = plan.UpdateSvcbHints.ValueBool()
	record.NameServer = plan.NameServer.ValueString()
	record.Cname = plan.Cname.ValueString()
	record.PtrName = plan.PtrName.ValueString()
	record.Exchange = plan.Exchange.ValueString()
	record.Preference = plan.Preference.ValueInt64()
	record.Text = plan.Text.ValueString()
	record.SplitText = plan.SplitText.ValueString()
	record.Protocol = plan.Protocol.ValueString()
	record.Forwarder = plan.Forwarder.ValueString()
	record.ForwarderPriority = plan.ForwarderPriority.ValueInt64()
	record.DnssecValidation = plan.DnssecValidation.ValueBool()
	record.ProxyType = plan.ProxyType.ValueString()
	record.ProxyAddress = plan.ProxyAddress.ValueString()
	record.ProxyPort = plan.ProxyPort.ValueInt64()
	record.ProxyUsername = plan.ProxyUsername.ValueString()
	record.ProxyPassword = plan.ProxyPassword.ValueString()

	err := r.client.CreateDnsZoneRecord(record, ctx)
	if err != nil {
		return err
	}

	return nil
}

// Read refreshes the Terraform state with the latest data.
func (r *dnsZoneRecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state dnsZoneRecordCreate
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if the record exists
	record, err := r.client.GetDnsZoneRecord(state.Domain.ValueString(), state.Type.ValueString(), ctx)
	if err != nil {

		tflog.Debug(ctx, "Removing record "+state.Domain.ValueString()+"from state due to error: "+err.Error())
		resp.State.RemoveResource(ctx)

		//resp.Diagnostics.AddError(
		//	"Error Reading DNS zone record",
		//	"Could not read DNS zone record  "+state.Domain.ValueString()+": "+err.Error(),
		//)
		return
	}

	state.Domain = types.StringValue(record.Name)
	state.Type = types.StringValue(record.Type)
	state.TTL = types.Int64Value(record.TTL)
	state.Cname = types.StringValue(record.RecordData.Cname)
	state.IPAddress = types.StringValue(record.RecordData.IpAddress)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dnsZoneRecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state dnsZoneRecordCreate
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	record := technitium.DnsZoneRecordCreate{}

	record.Type = state.Type.ValueString()
	record.Domain = state.Domain.ValueString()
	record.Zone = state.Zone.ValueString()
	record.IPAddress = state.IPAddress.ValueString()

	err := r.client.DeleteDnsZoneRecord(record, ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DNS zone record",
			"Could not delete DNS Zone record "+state.Domain.ValueString()+": "+err.Error(),
		)
		return
	}
}

func (r *dnsZoneRecordResource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	resp.RequiresReplace = path.Paths{
		path.Empty(),
	}

}
