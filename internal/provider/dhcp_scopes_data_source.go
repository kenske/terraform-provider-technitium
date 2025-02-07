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
	_ datasource.DataSource              = &dhcpScopesDataSource{}
	_ datasource.DataSourceWithConfigure = &dhcpScopesDataSource{}
)

func NewDhcpScopesDataSource() datasource.DataSource {
	return &dhcpScopesDataSource{}
}

type dhcpScopesDataSource struct {
	client *technitium.Client
}
type dhcpScopesDataSourceModel struct {
	Scopes []dhcpScopeModel `tfsdk:"scopes"`
}

type dhcpScopeModel struct {
	Name             types.String `tfsdk:"name"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	StartingAddress  types.String `tfsdk:"starting_address"`
	EndingAddress    types.String `tfsdk:"ending_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	NetworkAddress   types.String `tfsdk:"network_address"`
	BroadcastAddress types.String `tfsdk:"broadcast_address"`
	InterfaceAddress types.String `tfsdk:"interface_address"`
}

// Metadata returns the data source type name.
func (d *dhcpScopesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp_scopes"
}

// Schema defines the schema for the data source.
func (d *dhcpScopesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"scopes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"starting_address": schema.StringAttribute{
							Computed: true,
						},
						"ending_address": schema.StringAttribute{
							Computed: true,
						},
						"subnet_mask": schema.StringAttribute{
							Computed: true,
						},
						"network_address": schema.StringAttribute{
							Computed: true,
						},
						"broadcast_address": schema.StringAttribute{
							Computed: true,
						},
						"interface_address": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *dhcpScopesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state dhcpScopesDataSourceModel

	scopes, err := d.client.GetScopes()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read DHCP Scopes",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, scope := range scopes {
		scopeState := dhcpScopeModel{
			Name:             types.StringValue(scope.Name),
			Enabled:          types.BoolValue(scope.Enabled),
			StartingAddress:  types.StringValue(scope.StartingAddress),
			EndingAddress:    types.StringValue(scope.EndingAddress),
			SubnetMask:       types.StringValue(scope.SubnetMask),
			NetworkAddress:   types.StringValue(scope.NetworkAddress),
			BroadcastAddress: types.StringValue(scope.BroadcastAddress),
			InterfaceAddress: types.StringValue(scope.InterfaceAddress),
		}

		state.Scopes = append(state.Scopes, scopeState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *dhcpScopesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*technitium.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
