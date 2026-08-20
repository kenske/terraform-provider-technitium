package technitium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestClient_doRequest_TokenAuth(t *testing.T) {
	ctx := context.Background()

	t.Run("default mode sends Authorization Bearer header and no token query param", func(t *testing.T) {
		var gotAuthHeader, gotRawQuery string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotAuthHeader = r.Header.Get("Authorization")
			gotRawQuery = r.URL.RawQuery
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"data":"test"}`)
		}))
		defer server.Close()

		client := &Client{HostURL: server.URL, HTTPClient: server.Client(), Token: "test-token"}
		req, _ := http.NewRequest("GET", server.URL+"/api/test", nil)

		if _, err := client.doRequest(req, ctx); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if gotAuthHeader != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got %q", gotAuthHeader)
		}
		if strings.Contains(gotRawQuery, "token=") {
			t.Errorf("Expected no 'token' query parameter in request URL, got raw query %q", gotRawQuery)
		}
	})

	t.Run("legacy mode sends token as query parameter and no Authorization header", func(t *testing.T) {
		var gotAuthHeader, gotToken string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotAuthHeader = r.Header.Get("Authorization")
			gotToken = r.URL.Query().Get("token")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"data":"test"}`)
		}))
		defer server.Close()

		client := &Client{HostURL: server.URL, HTTPClient: server.Client(), Token: "test-token", LegacyTokenAuth: true}
		req, _ := http.NewRequest("GET", server.URL+"/api/test", nil)

		if _, err := client.doRequest(req, ctx); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if gotToken != "test-token" {
			t.Errorf("Expected 'token' query parameter 'test-token', got %q", gotToken)
		}
		if gotAuthHeader != "" {
			t.Errorf("Expected no Authorization header in legacy mode, got %q", gotAuthHeader)
		}
	})
}

func TestGetToken(t *testing.T) {
	t.Run("sends credentials as a POST form body, not a URL query string", func(t *testing.T) {
		var gotMethod, gotRawQuery, gotContentType string
		var gotUser, gotPass string

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotMethod = r.Method
			gotRawQuery = r.URL.RawQuery
			gotContentType = r.Header.Get("Content-Type")

			if err := r.ParseForm(); err != nil {
				t.Fatalf("failed to parse form body: %v", err)
			}
			gotUser = r.PostForm.Get("user")
			gotPass = r.PostForm.Get("pass")

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"token":"secret-session-token"}`)
		}))
		defer server.Close()

		// Password intentionally contains characters ('&', '=', '%', '#')
		// that would corrupt or hijack a hand-built query string if the
		// credentials were placed directly in the URL.
		const username = "admin"
		const password = `p&a=s%s#word`

		token, err := GetToken(server.URL, username, password)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if token != "secret-session-token" {
			t.Errorf("Expected token 'secret-session-token', got %q", token)
		}

		if gotMethod != http.MethodPost {
			t.Errorf("Expected POST request, got %q", gotMethod)
		}
		if gotRawQuery != "" {
			t.Errorf("Expected no credentials in the URL query string, got %q", gotRawQuery)
		}
		if gotContentType != "application/x-www-form-urlencoded" {
			t.Errorf("Expected form-urlencoded Content-Type, got %q", gotContentType)
		}
		if gotUser != username {
			t.Errorf("Expected form field user=%q, got %q", username, gotUser)
		}
		if gotPass != password {
			t.Errorf("Expected form field pass=%q, got %q", password, gotPass)
		}
	})

	t.Run("missing credentials", func(t *testing.T) {
		if _, err := GetToken("http://testhost", "", "password"); err == nil {
			t.Error("Expected error for missing username, got nil")
		}
		if _, err := GetToken("http://testhost", "user", ""); err == nil {
			t.Error("Expected error for missing password, got nil")
		}
	})
}

func TestClient_doRequest_TransportErrorRedactsToken(t *testing.T) {
	ctx := context.Background()

	t.Run("legacy mode transport failure does not leak the token", func(t *testing.T) {
		mockScenario := test.Scenario{
			ExpectedError: fmt.Errorf("simulated connection refused"),
		}
		client, cleanup := GetMockClient(mockScenario)
		defer cleanup()
		client.LegacyTokenAuth = true
		client.Token = "super-secret-live-token"

		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)

		_, err := client.doRequest(req, ctx)
		if err == nil {
			t.Fatal("Expected an error from the simulated transport failure, got nil")
		}

		if strings.Contains(err.Error(), client.Token) {
			t.Errorf("Expected the live token to be redacted from the error, got %v", err)
		}
		if !strings.Contains(err.Error(), "token=hidden") {
			t.Errorf("Expected the redacted error to contain 'token=hidden', got %v", err)
		}
		// The underlying failure reason must still be observable for debugging.
		if !strings.Contains(err.Error(), "simulated connection refused") {
			t.Errorf("Expected the underlying error reason to be preserved, got %v", err)
		}
	})
}
