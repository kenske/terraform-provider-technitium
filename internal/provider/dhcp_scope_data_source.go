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
	_ datasource.DataSource              = &dhcpScopeDataSource{}
	_ datasource.DataSourceWithConfigure = &dhcpScopeDataSource{}
)

func NewDhcpScopeDataSource() datasource.DataSource {
	return &dhcpScopeDataSource{}
}

type dhcpScopeDataSource struct {
	client *technitium.Client
}

// Metadata returns the data source type name.
func (d *dhcpScopeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp_scope"
}

// Schema defines the schema for the data source.
func (d *dhcpScopeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: DhcpScopeSchema(false),
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dhcpScopeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state dhcpScope
	diags := resp.State.Get(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	scope, err := d.client.GetScope(state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DHCP Scope",
			err.Error(),
		)
		return
	}

	// Map response body to model
	data := dhcpScope{
		Name:             types.StringValue(scope.Name),
		Enabled:          types.BoolValue(scope.Enabled),
		StartingAddress:  types.StringValue(scope.StartingAddress),
		EndingAddress:    types.StringValue(scope.EndingAddress),
		SubnetMask:       types.StringValue(scope.SubnetMask),
		NetworkAddress:   types.StringValue(scope.NetworkAddress),
		BroadcastAddress: types.StringValue(scope.BroadcastAddress),
		InterfaceAddress: types.StringValue(scope.InterfaceAddress),
	}

	// Set state
	diagsState := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *dhcpScopeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
