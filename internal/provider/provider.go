package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	"terraform-provider-technitium-dns/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &technitiumProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &technitiumProvider{
			version: version,
		}
	}
}

type technitiumProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type technitiumProviderModel struct {
	Host  types.String `tfsdk:"host"`
	Token types.String `tfsdk:"token"`
}

// Metadata returns the provider type name.
func (p *technitiumProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "technitium"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *technitiumProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *technitiumProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config technitiumProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Technitium DNS API Host",
			"The provider cannot create the Technitium DNS API client as there is an unknown configuration value for the host. ",
		)
	}

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown Technitium DNS API Password",
			"The provider cannot create the Technitium DNS API client as there is an unknown configuration value for the API token. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("TECHNITIUM_HOST")
	token := os.Getenv("TECHNITIUM_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Technitium DNS API API Host",
			"Host value is missing",
		)
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Technitium DNS API token",
			"Token is empty",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := technitium.NewClient(host, token, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Technitium DNS API Client",
			"An unexpected error occurred when creating the Technitium DNS API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	ctx = tflog.SetField(ctx, "technitium_api_host", host)
	tflog.Info(ctx, "Configured API client", map[string]any{"success": true})
}

func (p *technitiumProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDhcpScopesDataSource,
	}
}

func (p *technitiumProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDhcpScopeResource,
	}
}
