package technitium

import "terraform-provider-technitium/internal/test"

func GetMockClient(scenario test.Scenario) (*Client, func()) {

	hostURL, httpClient, cleanup := test.GetTestClientComponents(scenario)

	client := &Client{
		HostURL:    hostURL,
		HTTPClient: httpClient,
		Token:      "test-token",
	}

	return client, cleanup
}
