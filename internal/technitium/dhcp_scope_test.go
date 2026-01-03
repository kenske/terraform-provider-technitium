package technitium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-technitium/internal/test"
	"testing"
)

func TestClient_GetScopes(t *testing.T) {
	ctx := context.Background()

	t.Run("successful get scopes", func(t *testing.T) {
		mockResponse := DhcpScopesResponse{
			Response: struct {
				Scopes []DhcpScopeList
			}{
				Scopes: []DhcpScopeList{
					{
						Name:             "TestScope1",
						Enabled:          true,
						StartingAddress:  "192.168.1.100",
						EndingAddress:    "192.168.1.200",
						SubnetMask:       "255.255.255.0",
						NetworkAddress:   "192.168.1.0",
						BroadcastAddress: "192.168.1.255",
					},
					{
						Name:             "TestScope2",
						Enabled:          false,
						StartingAddress:  "10.0.0.10",
						EndingAddress:    "10.0.0.50",
						SubnetMask:       "255.255.255.0",
						NetworkAddress:   "10.0.0.0",
						BroadcastAddress: "10.0.0.255",
					},
				},
			},
			BaseResponse: BaseResponse{
				Status: "ok",
			},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scopes, err := client.GetScopes(ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(scopes) != 2 {
			t.Errorf("Expected 2 scopes, got %d", len(scopes))
		}

		if scopes[0].Name != "TestScope1" {
			t.Errorf("Expected first scope name 'TestScope1', got '%s'", scopes[0].Name)
		}

		if scopes[0].StartingAddress != "192.168.1.100" {
			t.Errorf("Expected starting address '192.168.1.100', got '%s'", scopes[0].StartingAddress)
		}

		if !scopes[0].Enabled {
			t.Error("Expected first scope to be enabled")
		}

		if scopes[1].Enabled {
			t.Error("Expected second scope to be disabled")
		}
	})

	t.Run("empty scopes list", func(t *testing.T) {
		mockResponse := DhcpScopesResponse{
			Response: struct {
				Scopes []DhcpScopeList
			}{
				Scopes: []DhcpScopeList{},
			},
			BaseResponse: BaseResponse{
				Status: "ok",
			},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scopes, err := client.GetScopes(ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(scopes) != 0 {
			t.Errorf("Expected 0 scopes, got %d", len(scopes))
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("network error"),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScopes(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "network error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("invalid json response", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `invalid json`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScopes(ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})

	t.Run("non-200 status code", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusInternalServerError,
			ExpectedBody:   `{"status":"error","errorMessage":"Internal server error"}`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScopes(ctx)
		if err == nil {
			t.Fatal("Expected error for non-200 status, got nil")
		}
	})
}

func TestClient_GetScope(t *testing.T) {
	ctx := context.Background()

	t.Run("successful get scope", func(t *testing.T) {
		mockScope := DhcpScope{
			Name:             "TestScope",
			StartingAddress:  "192.168.1.100",
			EndingAddress:    "192.168.1.200",
			SubnetMask:       "255.255.255.0",
			RouterAddress:    "192.168.1.1",
			UseThisDnsServer: true,
			DnsServers:       []string{"8.8.8.8", "8.8.4.4"},
			LeaseTimeDays:    7,
			LeaseTimeHours:   0,
			LeaseTimeMinutes: 0,
			DomainName:       "example.com",
			Exclusions: []Exclusion{
				{
					StartingAddress: "192.168.1.150",
					EndingAddress:   "192.168.1.160",
				},
			},
		}

		mockResponse := DhcpScopeResponse{
			Response:     mockScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.GetScope("TestScope", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if scope.Name != "TestScope" {
			t.Errorf("Expected name 'TestScope', got '%s'", scope.Name)
		}

		if scope.StartingAddress != "192.168.1.100" {
			t.Errorf("Expected starting address '192.168.1.100', got '%s'", scope.StartingAddress)
		}

		if scope.RouterAddress != "192.168.1.1" {
			t.Errorf("Expected router address '192.168.1.1', got '%s'", scope.RouterAddress)
		}

		if len(scope.DnsServers) != 2 {
			t.Errorf("Expected 2 DNS servers, got %d", len(scope.DnsServers))
		}

		if len(scope.Exclusions) != 1 {
			t.Errorf("Expected 1 exclusion, got %d", len(scope.Exclusions))
		}

		if scope.LeaseTimeDays != 7 {
			t.Errorf("Expected lease time days 7, got %d", scope.LeaseTimeDays)
		}
	})

	t.Run("scope not found", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   `{"status":"error","errorMessage":"Scope not found"}`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScope("NonExistentScope", ctx)
		if err == nil {
			t.Fatal("Expected error for non-existent scope, got nil")
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("connection timeout"),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScope("TestScope", ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "connection timeout") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("invalid json response", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{invalid json}`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetScope("TestScope", ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})
}

func TestClient_SetScope(t *testing.T) {
	ctx := context.Background()

	t.Run("successful create scope", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:             "NewScope",
			StartingAddress:  "192.168.2.100",
			EndingAddress:    "192.168.2.200",
			SubnetMask:       "255.255.255.0",
			RouterAddress:    "192.168.2.1",
			UseThisDnsServer: true,
			LeaseTimeDays:    1,
			LeaseTimeHours:   0,
			LeaseTimeMinutes: 0,
		}

		mockResponse := DhcpScopeResponse{
			Response:     inputScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.SetScope(inputScope, "", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if scope.Name != "NewScope" {
			t.Errorf("Expected name 'NewScope', got '%s'", scope.Name)
		}

		if scope.StartingAddress != "192.168.2.100" {
			t.Errorf("Expected starting address '192.168.2.100', got '%s'", scope.StartingAddress)
		}

		if scope.RouterAddress != "192.168.2.1" {
			t.Errorf("Expected router address '192.168.2.1', got '%s'", scope.RouterAddress)
		}
	})

	t.Run("successful update scope with rename", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:            "RenamedScope",
			StartingAddress: "192.168.3.100",
			EndingAddress:   "192.168.3.200",
			SubnetMask:      "255.255.255.0",
		}

		mockResponse := DhcpScopeResponse{
			Response:     inputScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.SetScope(inputScope, "OldScope", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if scope.Name != "RenamedScope" {
			t.Errorf("Expected name 'RenamedScope', got '%s'", scope.Name)
		}
	})

	t.Run("scope with multiple exclusions", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:            "MultiExclusionScope",
			StartingAddress: "192.168.10.1",
			EndingAddress:   "192.168.10.254",
			SubnetMask:      "255.255.255.0",
			RouterAddress:   "192.168.10.1",
			LeaseTimeDays:   7,
			Exclusions: []Exclusion{
				{
					StartingAddress: "192.168.10.1",
					EndingAddress:   "192.168.10.10",
				},
				{
					StartingAddress: "192.168.10.100",
					EndingAddress:   "192.168.10.110",
				},
				{
					StartingAddress: "192.168.10.200",
					EndingAddress:   "192.168.10.210",
				},
			},
		}

		mockResponse := DhcpScopeResponse{
			Response:     inputScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.SetScope(inputScope, "", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(scope.Exclusions) != 3 {
			t.Errorf("Expected 3 exclusions, got %d", len(scope.Exclusions))
		}

		if scope.Exclusions[0].StartingAddress != "192.168.10.1" {
			t.Errorf("Expected first exclusion starting address '192.168.10.1', got '%s'", scope.Exclusions[0].StartingAddress)
		}

		if scope.Exclusions[2].EndingAddress != "192.168.10.210" {
			t.Errorf("Expected third exclusion ending address '192.168.10.210', got '%s'", scope.Exclusions[2].EndingAddress)
		}
	})

	t.Run("scope with vendor info and tftp servers", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:                "VendorScope",
			StartingAddress:     "10.1.0.10",
			EndingAddress:       "10.1.0.100",
			SubnetMask:          "255.255.255.0",
			RouterAddress:       "10.1.0.1",
			TftpServerAddresses: []string{"10.1.0.5", "10.1.0.6"},
			BootFileName:        "pxelinux.0",
			NextServerAddress:   "10.1.0.5",
			VendorInfo: []DhcpVendorInfo{
				{
					Identifier:  "MSFT 5.0",
					Information: "Microsoft Windows",
				},
			},
			CAPWAPAcIpAddresses: []string{"10.1.0.10"},
			LeaseTimeDays:       1,
		}

		mockResponse := DhcpScopeResponse{
			Response:     inputScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.SetScope(inputScope, "", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(scope.VendorInfo) != 1 {
			t.Errorf("Expected 1 vendor info, got %d", len(scope.VendorInfo))
		}

		if scope.VendorInfo[0].Identifier != "MSFT 5.0" {
			t.Errorf("Expected vendor identifier 'MSFT 5.0', got '%s'", scope.VendorInfo[0].Identifier)
		}

		if len(scope.TftpServerAddresses) != 2 {
			t.Errorf("Expected 2 TFTP servers, got %d", len(scope.TftpServerAddresses))
		}

		if scope.BootFileName != "pxelinux.0" {
			t.Errorf("Expected boot file name 'pxelinux.0', got '%s'", scope.BootFileName)
		}
	})

	t.Run("scope with complex configuration", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:                         "ComplexScope",
			StartingAddress:              "10.0.0.10",
			EndingAddress:                "10.0.0.100",
			SubnetMask:                   "255.255.255.0",
			RouterAddress:                "10.0.0.1",
			UseThisDnsServer:             false,
			DnsServers:                   []string{"1.1.1.1", "1.0.0.1"},
			WinsServers:                  []string{"10.0.0.5"},
			NtpServers:                   []string{"time.google.com"},
			DomainName:                   "test.local",
			DomainSearchList:             []string{"test.local", "example.com"},
			LeaseTimeDays:                7,
			LeaseTimeHours:               12,
			LeaseTimeMinutes:             30,
			OfferDelayTime:               100,
			PingCheckEnabled:             true,
			PingCheckTimeout:             1000,
			PingCheckRetries:             3,
			AllowOnlyReservedLeases:      false,
			IgnoreClientIdentifierOption: true,
			StaticRoutes: []DhcpStaticRoute{
				{
					Destination: "172.16.0.0",
					SubnetMask:  "255.255.0.0",
					Router:      "10.0.0.254",
				},
			},
			Exclusions: []Exclusion{
				{
					StartingAddress: "10.0.0.50",
					EndingAddress:   "10.0.0.60",
				},
			},
			GenericOptions: []DhcpGenericOption{
				{
					Code:  252,
					Value: "http://wpad.example.com/wpad.dat",
				},
			},
		}

		mockResponse := DhcpScopeResponse{
			Response:     inputScope,
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		scope, err := client.SetScope(inputScope, "", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(scope.StaticRoutes) != 1 {
			t.Errorf("Expected 1 static route, got %d", len(scope.StaticRoutes))
		}

		if len(scope.Exclusions) != 1 {
			t.Errorf("Expected 1 exclusion, got %d", len(scope.Exclusions))
		}

		if len(scope.GenericOptions) != 1 {
			t.Errorf("Expected 1 generic option, got %d", len(scope.GenericOptions))
		}

		if scope.GenericOptions[0].Code != 252 {
			t.Errorf("Expected generic option code 252, got %d", scope.GenericOptions[0].Code)
		}

		if scope.PingCheckEnabled != true {
			t.Error("Expected ping check to be enabled")
		}

		if scope.PingCheckTimeout != 1000 {
			t.Errorf("Expected ping check timeout 1000, got %d", scope.PingCheckTimeout)
		}

		if scope.IgnoreClientIdentifierOption != true {
			t.Error("Expected ignoreClientIdentifierOption to be true")
		}
	})

	t.Run("api returns error status", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:            "FailScope",
			StartingAddress: "192.168.1.100",
			EndingAddress:   "192.168.1.200",
			SubnetMask:      "255.255.255.0",
		}

		mockResponse := DhcpScopeResponse{
			BaseResponse: BaseResponse{
				Status:       "error",
				ErrorMessage: "Invalid subnet configuration",
			},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.SetScope(inputScope, "", ctx)
		if err == nil {
			t.Fatal("Expected error for failed scope creation, got nil")
		}
		if !strings.Contains(err.Error(), "failed to create scope") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("http client error", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:            "TestScope",
			StartingAddress: "192.168.1.100",
			EndingAddress:   "192.168.1.200",
			SubnetMask:      "255.255.255.0",
		}

		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("network unreachable"),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.SetScope(inputScope, "", ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "network unreachable") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("invalid json response", func(t *testing.T) {
		inputScope := DhcpScope{
			Name:            "TestScope",
			StartingAddress: "192.168.1.100",
			EndingAddress:   "192.168.1.200",
			SubnetMask:      "255.255.255.0",
		}

		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `not valid json`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.SetScope(inputScope, "", ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})
}

func TestClient_DeleteScope(t *testing.T) {
	ctx := context.Background()

	t.Run("successful delete scope", func(t *testing.T) {
		mockResponse := DhcpScopeResponse{
			BaseResponse: BaseResponse{Status: "ok"},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteScope("TestScope", ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("delete non-existent scope", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   `{"status":"error","errorMessage":"Scope not found"}`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteScope("NonExistentScope", ctx)
		if err == nil {
			t.Fatal("Expected error for non-existent scope, got nil")
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("connection refused"),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteScope("TestScope", ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "connection refused") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("invalid json response", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{bad json`,
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteScope("TestScope", ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})

	t.Run("api returns error status", func(t *testing.T) {
		mockResponse := DhcpScopeResponse{
			BaseResponse: BaseResponse{
				Status:       "error",
				ErrorMessage: "Cannot delete scope: active leases exist",
			},
		}

		mockBody, _ := json.Marshal(mockResponse)
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   string(mockBody),
		}

		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteScope("TestScope", ctx)
		if err == nil {
			t.Fatal("Expected error for scope with active leases, got nil")
		}
	})
}
