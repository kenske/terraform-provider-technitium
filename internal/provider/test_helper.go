package provider

import (
	"terraform-provider-technitium/internal/technitium"
	"terraform-provider-technitium/internal/test"
)

func GetMockClient(scenario test.Scenario) (*technitium.Client, func()) {

	hostURL, httpClient, cleanup := test.GetTestClientComponents(scenario)

	client := &technitium.Client{
		HostURL:    hostURL,
		HTTPClient: httpClient,
		Token:      "test-token",
	}

	return client, cleanup
}
