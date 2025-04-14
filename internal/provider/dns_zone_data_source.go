package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-technitium-dns/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &dnsZoneDataSource{}
	_ datasource.DataSourceWithConfigure = &dnsZoneDataSource{}
)

func NewDnsZoneDataSource() datasource.DataSource {
	return &dnsZoneDataSource{}
}

type dnsZoneDataSource struct {
	client *technitium.Client
}

// Metadata returns the data source type name.
func (d *dnsZoneDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone"
}

// Schema defines the schema for the data source.
func (d *dnsZoneDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DnsZoneSchema(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dnsZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state dnsZoneGet
	diags := resp.State.Get(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, err := d.client.GetDnsZone(state.Name.ValueString(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DNS Zone",
			err.Error(),
		)
		return
	}

	// Map response body to model
	data := dnsZoneGet{}
	data.Name = types.StringValue(zone.Name)
	data.Type = types.StringValue(zone.Type)
	data.Catalog = types.StringValue(zone.Catalog)
	data.Disabled = types.BoolValue(zone.Disabled)
	data.DnssecStatus = types.StringValue(zone.DnsSecStatus)
	data.NotifyFailed = types.BoolValue(zone.NotifyFailed)
	data.NotifyFailedFor = convertStringListToTF(zone.NotifyFailedFor)
	data.QueryAccess = types.StringValue(zone.QueryAccess)
	data.QueryAccessNetworkAcl = convertStringListToTF(zone.QueryAccessNetworkAcl)
	data.ZoneTransfer = types.StringValue(zone.ZoneTransfer)
	data.ZoneTransferNetworkAcl = convertStringListToTF(zone.ZoneTransferNetworkAcl)
	data.ZoneTransferTsigKeyNames = convertStringListToTF(zone.ZoneTransferTsigKeyNames)
	data.Notify = types.StringValue(zone.Notify)
	data.NotifyNameServers = convertStringListToTF(zone.NotifyNameServers)
	data.Update = types.StringValue(zone.Update)
	data.UpdateNetworkAcl = convertStringListToTF(zone.UpdateNetworkAcl)

	// Set state
	diagsState := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *dnsZoneDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}
