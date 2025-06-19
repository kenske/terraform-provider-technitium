package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-technitium/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &dnsZoneRecordDataSource{}
	_ datasource.DataSourceWithConfigure = &dnsZoneRecordDataSource{}
)

func NewDnsZoneRecordDataSource() datasource.DataSource {
	return &dnsZoneRecordDataSource{}
}

type dnsZoneRecordDataSource struct {
	client *technitium.Client
}

// Metadata returns the data source type name.
func (d *dnsZoneRecordDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone_record"
}

// Schema defines the schema for the data source.
func (d *dnsZoneRecordDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DnsZoneRecordDataSourceSchema(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dnsZoneRecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state dnsZoneRecord
	diags := resp.State.Get(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	record, err := d.client.GetDnsZoneRecord(state.Name.ValueString(), state.Type.ValueString(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DNS Zone",
			err.Error(),
		)
		return
	}

	ZoneRecord := parseRecord(record)

	// Set state
	diagsState := resp.State.Set(ctx, &ZoneRecord)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *dnsZoneRecordDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = ConfigureDataSourceClient(req, resp)
}
