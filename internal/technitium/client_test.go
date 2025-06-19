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

func TestGetSessionInfo(t *testing.T) {
	ctx := context.Background()

	t.Run("successful session info", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"success", "message":"Session is valid."}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.GetSessionInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("missing token", func(t *testing.T) {

		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"status":"success", "message":"Session is valid."}`,
		}

		server := test.NewTestServer(mockScenario)
		defer server.Close()

		client := &Client{HostURL: "http://testhost", HTTPClient: server.Client(), Token: ""}
		err := client.GetSessionInfo(ctx)
		if err == nil {
			t.Errorf("Expected error for missing token, got nil")
		}
		if !strings.Contains(err.Error(), "missing API token") {
			t.Errorf("Expected 'missing API token' error, got %v", err)
		}
	})

	t.Run("http client error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("simulated HTTP client error"),
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.GetSessionInfo(ctx)
		if err == nil {
			t.Errorf("Expected HTTP client error, got nil")
		}
		if !strings.Contains(err.Error(), "simulated HTTP client error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("non-200 status code", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedBody:   `{"status":"error", "message":"Invalid token"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.GetSessionInfo(ctx)
		if err == nil {
			t.Errorf("Expected error for non-200 status, got nil")
		}
		expectedErrMsg := fmt.Sprintf("status: %d, body: %s", http.StatusUnauthorized, mockScenario.ExpectedBody)
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Expected error message '%s', got '%v'", expectedErrMsg, err)
		}
	})

	t.Run("json unmarshal error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `invalid json`, // This will cause unmarshal to fail for BaseResponse
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		err := client.GetSessionInfo(ctx)
		if err == nil {
			t.Errorf("Expected JSON unmarshal error, got nil")
		}
		// Check for a generic JSON error, as the exact message can vary.
		// This test assumes BaseResponse expects a certain structure.
		// If GetSessionInfo doesn't unmarshal, this test is different.
		// Current GetSessionInfo unmarshals into BaseResponse.
		var syntaxError *json.SyntaxError
		if !errors.As(err, &syntaxError) {
			t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
		}
	})
}

func TestClient_doRequest(t *testing.T) {
	ctx := context.Background()

	t.Run("successful request", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data":"test"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)

		body, err := client.doRequest(req, ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if string(body) != mockScenario.ExpectedBody {
			t.Errorf("Expected body '%s', got '%s'", mockScenario.ExpectedBody, string(body))
		}
	})

	t.Run("http client Do error", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("simulated Do error"),
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()

		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)

		_, err := client.doRequest(req, ctx)
		if err == nil {
			t.Errorf("Expected error from HTTPClient.Do, got nil")
		}
		if !strings.Contains(err.Error(), "simulated Do error") {
			t.Errorf("Error message mismatch, got %v", err)
		}
	})

	t.Run("non-200 status code", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `{"error":"bad request"}`,
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()
		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)

		_, err := client.doRequest(req, ctx)
		if err == nil {
			t.Errorf("Expected error for non-200 status, got nil")
		}
		expectedErrMsg := fmt.Sprintf("status: %d, body: %s", http.StatusBadRequest, mockScenario.ExpectedBody)
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Expected error message '%s', got '%v'", expectedErrMsg, err)
		}
	})
}
