package provider

import (
	"context"
	"fmt"
	"terraform-provider-technitium-dns/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &dnsZonesDataSource{}
	_ datasource.DataSourceWithConfigure = &dnsZonesDataSource{}
)

func NewDnsZonesDataSource() datasource.DataSource {
	return &dnsZonesDataSource{}
}

type dnsZonesDataSource struct {
	client *technitium.Client
}
type dnsZonesDataSourceModel struct {
	Zones []dnsZoneList `tfsdk:"zones"`
}

// Metadata returns the data source type name.
func (d *dnsZonesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zones"
}

// Schema defines the schema for the data source.
func (d *dnsZonesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zones": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: DnsZonesSchema(),
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dnsZonesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dnsZonesDataSourceModel

	zones, err := d.client.GetDnsZones(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DNS Zones",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, zone := range zones {
		zoneState := dnsZoneList{}
		zoneState.Name = types.StringValue(zone.Name)
		zoneState.Type = types.StringValue(zone.Type)
		zoneState.Catalog = types.StringValue(zone.Catalog)
		zoneState.Disabled = types.BoolValue(zone.Disabled)
		zoneState.DnssecStatus = types.StringValue(zone.DnssecStatus)
		zoneState.SoaSerial = types.Int32Value(int32(zone.SoaSerial))
		zoneState.Expiry = types.StringValue(zone.Expiry)
		zoneState.IsExpired = types.BoolValue(zone.IsExpired)
		zoneState.LastModified = types.StringValue(zone.LastModified)
		zoneState.Internal = types.BoolValue(zone.Internal)

		state.Zones = append(state.Zones, zoneState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *dnsZonesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
