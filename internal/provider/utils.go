package provider

import (
	"context"
	"fmt"
	"terraform-provider-technitium/internal/technitium"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func setStringIfNotEmpty(target *types.String, value string) {
	if value != "" {
		*target = types.StringValue(value)
	}
}

func updateZoneRecord(ctx context.Context) error {
	host := "http://localhost:5380"
	username := "admin"
	password := "password"

	token, err := technitium.GetToken(host, username, password)

	if err != nil {
		return fmt.Errorf("Error getting token: %v\n", err)
	}

	client, err := technitium.NewClient(host, token, ctx)

	if err != nil {
		return fmt.Errorf("Error creating client: %v\n", err)
	}

	update := technitium.DnsZoneRecordUpdate{
		DnsZoneRecordCreate: technitium.DnsZoneRecordCreate{
			Domain:    "test.example.com",
			Type:      "A",
			Zone:      "example.com",
			IPAddress: "192.168.1.10",
			Comments:  "external update",
		},
	}

	err = client.UpdateDnsZoneRecord(update, ctx)
	if err != nil {
		return fmt.Errorf("Error updating DNS zone record: %v\n", err)
	}

	return nil

}
