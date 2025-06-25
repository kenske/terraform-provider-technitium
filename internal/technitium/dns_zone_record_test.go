package technitium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"terraform-provider-technitium/internal/test"
	"testing"
)

func TestClient_GetDnsZoneRecords(t *testing.T) {
	ctx := context.Background()
	domain := "example.com"

	t.Run("successful retrieval", func(t *testing.T) {
		expectedRecords := []DnsZoneRecord{
			{Name: "example.com", Type: "A", RecordData: DnsZoneRecordData{IpAddress: "192.0.2.1"}},
			{Name: "example.com", Type: "SOA", RecordData: DnsZoneRecordData{PrimaryNameServer: "dns-server"}},
		}
		mockResponseBody := `{"response":{"records":[{"name":"example.com","type":"A","rdata":{"ipAddress":"192.0.2.1"}},{"name":"example.com","type":"SOA","rdata":{"primaryNameServer":"dns-server"}}]}}`
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   mockResponseBody,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		records, err := client.GetDnsZoneRecords(domain, ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(records, expectedRecords) {
			t.Errorf("Expected records %+v, got %+v", expectedRecords, records)
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{ExpectedError: fmt.Errorf("simulated HTTP client error")}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecords(domain, ctx)
		if err == nil {
			t.Fatal("Expected HTTP client error, got nil")
		}
		if !strings.Contains(err.Error(), "simulated HTTP client error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("non-200 status code", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   `{"error":"not found"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecords(domain, ctx)
		if err == nil {
			t.Fatal("Expected error for non-200 status, got nil")
		}
		expectedErrMsg := fmt.Sprintf("status: %d, body: %s", http.StatusNotFound, mockScenario.ExpectedBody)
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Expected error message containing '%s', got '%v'", expectedErrMsg, err)
		}
	})

	t.Run("json unmarshal error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `invalid json`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecords(domain, ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})
}

func TestClient_GetDnsZoneRecord(t *testing.T) {
	ctx := context.Background()
	domain := "example.com"

	t.Run("successful retrieval of specific record", func(t *testing.T) {
		recordType := "SOA"
		expectedRecord := DnsZoneRecord{
			Name:         "example.com",
			Type:         "SOA",
			TTL:          0,
			Disabled:     false,
			DnsSecStatus: "Unknown",
			LastUsedOn:   "0001-01-01T00:00:00",
			LastModified: "2025-05-29T21:52:19.4867037Z",
			ExpiryTTL:    0,
			RecordData: DnsZoneRecordData{
				PrimaryNameServer:   "dns-server",
				ResponsiblePerson:   "invalid",
				Serial:              4,
				Refresh:             900,
				Retry:               300,
				Expire:              604800,
				Minimum:             900,
				UseSerialDateScheme: false,
			},
		}

		scenario := test.GetMockScenarioFromFile(t, "../test/mocks/dns_zone_records_response.json", http.StatusOK)
		client, cleanup := GetMockClient(scenario)
		defer cleanup()

		record, err := client.GetDnsZoneRecord(domain, recordType, ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !reflect.DeepEqual(record, expectedRecord) {
			test.PrintPrettyDeepEqualError(t, expectedRecord, record)
		}
	})

	t.Run("error from GetDnsZoneRecords", func(t *testing.T) {
		mockScenario := test.Scenario{ExpectedError: fmt.Errorf("simulated GetDnsZoneRecords error")}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecord(domain, "A", ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "failed to get DNS zone record") || !strings.Contains(err.Error(), "simulated GetDnsZoneRecords error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("no records found for domain", func(t *testing.T) {
		mockResponseBody := `{"response":{"records":[]}}`
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   mockResponseBody,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecord(domain, "A", ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), fmt.Sprintf("no DNS zone records found for domain: %s", domain)) {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("record of specific type not found", func(t *testing.T) {
		nonExistentType := "TXT"
		mockResponseBody := `{"response":{"records":[{"name":"example.com","type":"A","ipAddress":"192.0.2.1"}]}}`
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   mockResponseBody,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		_, err := client.GetDnsZoneRecord(domain, nonExistentType, ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		expectedErrMsg := fmt.Sprintf("no DNS zone record found for domain: %s with type: %s", domain, nonExistentType)
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Expected error message '%s', got '%v'", expectedErrMsg, err)
		}
	})
}

func TestClient_CreateDnsZoneRecord(t *testing.T) {
	ctx := context.Background()
	recordCreate := DnsZoneRecordCreate{Domain: "new.example.com", Type: "A", Zone: "example.com", IPAddress: "192.0.2.10"}

	t.Run("successful creation", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"ok"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.CreateDnsZoneRecord(recordCreate, ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("api error on creation", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK, // API itself might return 200 but with an error status in body
			ExpectedBody:   `{"status":"error", "errorMessage":"failed to create"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.CreateDnsZoneRecord(recordCreate, ctx)
		if err == nil {
			t.Fatal("Expected API error, got nil")
		}
		if !strings.Contains(err.Error(), "failed to create zone record: failed to create") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{ExpectedError: fmt.Errorf("simulated HTTP client error")}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.CreateDnsZoneRecord(recordCreate, ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "simulated HTTP client error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("json unmarshal error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `invalid json`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.CreateDnsZoneRecord(recordCreate, ctx)
		if err == nil {
			t.Fatal("Expected JSON unmarshal error, got nil")
		}
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})
}

func TestClient_UpdateDnsZoneRecord(t *testing.T) {
	ctx := context.Background()
	recordUpdate := DnsZoneRecordUpdate{
		DnsZoneRecordCreate: DnsZoneRecordCreate{
			Domain:    "update.example.com",
			Type:      "A",
			Zone:      "example.com",
			IPAddress: "192.0.2.11",
		},
	}

	t.Run("successful update", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"ok"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.UpdateDnsZoneRecord(recordUpdate, ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("api error on update", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"error", "errorMessage":"failed to update"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.UpdateDnsZoneRecord(recordUpdate, ctx)
		if err == nil {
			t.Fatal("Expected API error, got nil")
		}
		if !strings.Contains(err.Error(), "failed to update zone record: failed to update") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})
}

func TestClient_DeleteDnsZoneRecord(t *testing.T) {
	ctx := context.Background()
	recordDelete := DnsZoneRecordCreate{Domain: "delete.example.com", Type: "A", Zone: "example.com", IPAddress: "192.0.2.12"}

	t.Run("successful deletion", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"ok"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteDnsZoneRecord(recordDelete, ctx)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("deletion of non-existent record (treated as success)", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK, // API might return 200 OK
			ExpectedBody:   `{"status":"error", "errorMessage":"Cannot delete record: no such record exists"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteDnsZoneRecord(recordDelete, ctx)
		if err != nil {
			t.Fatalf("Expected no error for 'no such record exists', got %v", err)
		}
	})

	t.Run("api error on deletion (other than no such record)", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"error", "errorMessage":"some other delete error"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.DeleteDnsZoneRecord(recordDelete, ctx)
		if err == nil {
			t.Fatal("Expected API error, got nil")
		}
		if !strings.Contains(err.Error(), "failed to delete zone record: some other delete error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})
}
