package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-technitium/internal/technitium"
)

func convertStringListToTF(items []string) []types.String {

	if len(items) == 0 {
		return nil
	}

	var attr []types.String
	for _, item := range items {
		attr = append(attr, types.StringValue(item))
	}

	return attr

}

func convertTfListToStringList(items []types.String) []string {
	if len(items) == 0 {
		return nil
	}

	var attr []string
	for _, item := range items {
		attr = append(attr, item.ValueString())
	}

	return attr
}

func ConfigureResourceClient(req resource.ConfigureRequest, resp *resource.ConfigureResponse) *technitium.Client {
	if req.ProviderData == nil {
		return nil
	}

	client, ok := req.ProviderData.(*technitium.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Configure Type",
			fmt.Sprintf("Expected *technitium.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}

	return client
}

func ConfigureDataSourceClient(req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) *technitium.Client {
	if req.ProviderData == nil {
		return nil
	}

	client, ok := req.ProviderData.(*technitium.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Configure Type",
			fmt.Sprintf("Expected *technitium.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}

	return client
}
