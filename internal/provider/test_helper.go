package provider

import (
	"fmt"
	"os"
	"terraform-provider-technitium/internal/technitium"
	"terraform-provider-technitium/internal/test"
	"testing"
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

func GetFileConfig(t *testing.T, file string) string {

	providerBytes, err := os.ReadFile("testdata/provider.tf")
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	configBytes, err := os.ReadFile(fmt.Sprintf("testdata/%s", file))
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	return string(providerBytes) + string(configBytes)
}
