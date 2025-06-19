package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-technitium/internal/technitium"

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
		data.Records = append(data.Records, *parseRecord(record))

	}

	// Set state
	diagsState := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func parseRecord(record technitium.DnsZoneRecord) *dnsZoneRecord {
	return &dnsZoneRecord{
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
}

func getRecordData(record technitium.DnsZoneRecord) *dnsZoneRecordData {

	rd := record.RecordData
	data := dnsZoneRecordData{}

	if rd.PrimaryNameServer != "" {
		data.PrimaryNameServer = types.StringValue(rd.PrimaryNameServer)
	}
	if rd.ResponsiblePerson != "" {
		data.ResponsiblePerson = types.StringValue(rd.ResponsiblePerson)
	}
	if rd.Serial != 0 {
		data.Serial = types.Int64Value(rd.Serial)
	}
	if rd.Refresh != 0 {
		data.Refresh = types.Int64Value(rd.Refresh)
	}
	if rd.Retry != 0 {
		data.Retry = types.Int64Value(rd.Retry)
	}
	if rd.Expire != 0 {
		data.Expire = types.Int64Value(rd.Expire)
	}
	if rd.Minimum != 0 {
		data.Minimum = types.Int64Value(rd.Minimum)
	}
	if rd.Protocol != "" {
		data.Protocol = types.StringValue(rd.Protocol)
	}
	if rd.Forwarder != "" {
		data.Forwarder = types.StringValue(rd.Forwarder)
	}
	if rd.Priority != 0 {
		data.Priority = types.Int64Value(rd.Priority)
	}
	if rd.ProxyType != "" {
		data.ProxyType = types.StringValue(rd.ProxyType)
	}
	if rd.IpAddress != "" {
		data.IpAddress = types.StringValue(rd.IpAddress)
	}
	if rd.Cname != "" {
		data.Cname = types.StringValue(rd.Cname)
	}

	return &data
}

func (d *dnsZoneRecordsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = ConfigureDataSourceClient(req, resp)
}
