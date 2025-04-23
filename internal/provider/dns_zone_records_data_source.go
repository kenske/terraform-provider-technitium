package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-technitium-dns/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &dnsZoneRecordsDataSource{}
	_ datasource.DataSourceWithConfigure = &dnsZoneRecordsDataSource{}
)

func NewDnsZoneRecordsDataSource() datasource.DataSource {
	return &dnsZoneRecordsDataSource{}
}

type dnsZoneRecordsDataSource struct {
	client *technitium.Client
}

// Metadata returns the data source type name.
func (d *dnsZoneRecordsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone_records"
}

// Schema defines the schema for the data source.
func (d *dnsZoneRecordsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DnsZoneRecordsSchema(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dnsZoneRecordsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state dnsZoneRecords
	diags := resp.State.Get(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	records, err := d.client.GetDnsZoneRecords(state.Domain.ValueString(), ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DNS Zone",
			err.Error(),
		)
		return
	}

	data := dnsZoneRecords{}

	// Map response body to model
	for _, record := range records {

		ZoneRecord := dnsZoneRecord{
			Name:         types.StringValue(record.Name),
			Type:         types.StringValue(record.Type),
			TTL:          types.Int64Value(record.TTL),
			Disabled:     types.BoolValue(record.Disabled),
			DnssecStatus: types.StringValue(record.DnsSecStatus),
			LastUsedOn:   types.StringValue(record.LastUsedOn),
			LastModified: types.StringValue(record.LastModified),
			ExpiryTTL:    types.Int64Value(record.ExpiryTTL),
			RecordData:   getRecordData(record),
		}

		data.Records = append(data.Records, ZoneRecord)

	}

	// Set state
	diagsState := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getRecordData(record technitium.DnsZoneRecord) dnsZoneRecordData {

	RecordData := dnsZoneRecordData{
		PrimaryNameServer:   types.StringValue(record.RecordData.PrimaryNameServer),
		ResponsiblePerson:   types.StringValue(record.RecordData.ResponsiblePerson),
		Serial:              types.Int64Value(record.RecordData.Serial),
		Refresh:             types.Int64Value(record.RecordData.Refresh),
		Retry:               types.Int64Value(record.RecordData.Retry),
		Expire:              types.Int64Value(record.RecordData.Expire),
		Minimum:             types.Int64Value(record.RecordData.Minimum),
		UseSerialDateScheme: types.BoolValue(record.RecordData.UseSerialDateScheme),
		Protocol:            types.StringValue(record.RecordData.Protocol),
		Forwarder:           types.StringValue(record.RecordData.Forwarder),
		Priority:            types.Int64Value(record.RecordData.Priority),
		DnssecValidation:    types.BoolValue(record.RecordData.DnssecValidation),
		ProxyType:           types.StringValue(record.RecordData.ProxyType),
		IpAddress:           types.StringValue(record.RecordData.IpAddress),
		Cname:               types.StringValue(record.RecordData.Cname),
	}

	return RecordData
}

func (d *dnsZoneRecordsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = ConfigureDataSourceClient(req, resp)
}
