package technitium

import (
	"context"
	"net/http"
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

	//t.Run("missing token", func(t *testing.T) {
	//	client := &Client{HostURL: "http://testhost", HTTPClient: &test.MockHttpClient{}, Token: ""}
	//	err := client.GetSessionInfo(ctx)
	//	if err == nil {
	//		t.Errorf("Expected error for missing token, got nil")
	//	}
	//	if !strings.Contains(err.Error(), "missing API token") {
	//		t.Errorf("Expected 'missing API token' error, got %v", err)
	//	}
	//})
	//
	//t.Run("http client error", func(t *testing.T) {
	//	mockScenario := test.Scenario{
	//		ExpectedError: fmt.Errorf("simulated HTTP client error"),
	//	}
	//	mockHttp := test.GetMockClient(mockScenario)
	//	client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
	//
	//	err := client.GetSessionInfo(ctx)
	//	if err == nil {
	//		t.Errorf("Expected HTTP client error, got nil")
	//	}
	//	if !strings.Contains(err.Error(), "simulated HTTP client error") {
	//		t.Errorf("Error message mismatch, got %v", err)
	//	}
	//})
	//
	//t.Run("non-200 status code", func(t *testing.T) {
	//	mockScenario := test.Scenario{
	//		ExpectedStatus: http.StatusUnauthorized,
	//		ExpectedBody:   `{"status":"error", "message":"Invalid token"}`,
	//	}
	//	mockHttp := test.GetMockClient(mockScenario)
	//	client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
	//
	//	err := client.GetSessionInfo(ctx)
	//	if err == nil {
	//		t.Errorf("Expected error for non-200 status, got nil")
	//	}
	//	expectedErrMsg := fmt.Sprintf("status: %d, body: %s", http.StatusUnauthorized, mockScenario.ExpectedBody)
	//	if !strings.Contains(err.Error(), expectedErrMsg) {
	//		t.Errorf("Expected error message '%s', got '%v'", expectedErrMsg, err)
	//	}
	//})
	//
	//t.Run("json unmarshal error", func(t *testing.T) {
	//	mockScenario := test.Scenario{
	//		ExpectedStatus: http.StatusOK,
	//		ExpectedBody:   `invalid json`, // This will cause unmarshal to fail for BaseResponse
	//	}
	//	mockHttp := test.GetMockClient(mockScenario)
	//	client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
	//
	//	err := client.GetSessionInfo(ctx)
	//	if err == nil {
	//		t.Errorf("Expected JSON unmarshal error, got nil")
	//	}
	//	// Check for a generic JSON error, as the exact message can vary.
	//	// This test assumes BaseResponse expects a certain structure.
	//	// If GetSessionInfo doesn't unmarshal, this test is different.
	//	// Current GetSessionInfo unmarshals into BaseResponse.
	//	var syntaxError *json.SyntaxError
	//	if !errors.As(err, &syntaxError) {
	//		t.Errorf("Expected a json.SyntaxError, got %T: %v", err, err)
	//	}
	//})
}

//func TestClient_doRequest(t *testing.T) {
//	ctx := context.Background()
//
//	t.Run("successful request", func(t *testing.T) {
//		mockScenario := test.Scenario{
//			ExpectedStatus: http.StatusOK,
//			ExpectedBody:   `{"data":"test"}`,
//		}
//		mockHttp := test.GetMockClient(mockScenario)
//		client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
//		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)
//
//		body, err := client.doRequest(req, ctx)
//		if err != nil {
//			t.Errorf("Expected no error, got %v", err)
//		}
//		if string(body) != mockScenario.ExpectedBody {
//			t.Errorf("Expected body '%s', got '%s'", mockScenario.ExpectedBody, string(body))
//		}
//	})
//
//	t.Run("http client Do error", func(t *testing.T) {
//		mockScenario := test.Scenario{
//			ExpectedError: fmt.Errorf("simulated Do error"),
//		}
//		mockHttp := test.GetMockClient(mockScenario)
//		client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
//		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)
//
//		_, err := client.doRequest(req, ctx)
//		if err == nil {
//			t.Errorf("Expected error from HTTPClient.Do, got nil")
//		}
//		if !strings.Contains(err.Error(), "simulated Do error") {
//			t.Errorf("Error message mismatch, got %v", err)
//		}
//	})
//
//	t.Run("read body error", func(t *testing.T) {
//		// Simulate an error during io.ReadAll
//		errorReader := &errorReader{readErr: fmt.Errorf("simulated read error")}
//		mockHttp := &test.MockHttpClient{
//			DoFunc: func(req *http.Request) (*http.Response, error) {
//				return &http.Response{
//					StatusCode: http.StatusOK,
//					Body:       io.NopCloser(errorReader),
//					Header:     make(http.Header),
//				}, nil
//			},
//		}
//		client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
//		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)
//
//		_, err := client.doRequest(req, ctx)
//		if err == nil {
//			t.Errorf("Expected error from io.ReadAll, got nil")
//		}
//		if !strings.Contains(err.Error(), "simulated read error") {
//			t.Errorf("Error message mismatch, got %v", err)
//		}
//	})
//
//	t.Run("non-200 status code", func(t *testing.T) {
//		mockScenario := test.Scenario{
//			ExpectedStatus: http.StatusBadRequest,
//			ExpectedBody:   `{"error":"bad request"}`,
//		}
//		mockHttp := test.GetMockClient(mockScenario)
//		client := &Client{HostURL: "http://testhost", HTTPClient: mockHttp, Token: "test-token"}
//		req, _ := http.NewRequest("GET", client.HostURL+"/api/test", nil)
//
//		_, err := client.doRequest(req, ctx)
//		if err == nil {
//			t.Errorf("Expected error for non-200 status, got nil")
//		}
//		expectedErrMsg := fmt.Sprintf("status: %d, body: %s", http.StatusBadRequest, mockScenario.ExpectedBody)
//		if !strings.Contains(err.Error(), expectedErrMsg) {
//			t.Errorf("Expected error message '%s', got '%v'", expectedErrMsg, err)
//		}
//	})
//}

// Helper for simulating read errors
type errorReader struct {
	readErr error
}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, er.readErr
}

//func TestClient_GetRequest(t *testing.T) {
//	client := &Client{HostURL: "http://example.com"}
//
//	t.Run("valid path", func(t *testing.T) {
//		path := "/api/data"
//		req, err := client.GetRequest(path)
//		if err != nil {
//			t.Fatalf("GetRequest failed: %v", err)
//		}
//		if req.Method != "GET" {
//			t.Errorf("Expected GET method, got %s", req.Method)
//		}
//		expectedURL := "http://example.com/api/data"
//		if req.URL.String() != expectedURL {
//			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
//		}
//	})
//
//	t.Run("empty path", func(t *testing.T) {
//		path := ""
//		req, err := client.GetRequest(path)
//		if err != nil {
//			t.Fatalf("GetRequest failed: %v", err)
//		}
//		expectedURL := "http://example.com"
//		if req.URL.String() != expectedURL {
//			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
//		}
//	})
//
//	t.Run("path with leading slash", func(t *testing.T) {
//		path := "/test"
//		req, err := client.GetRequest(path)
//		if err != nil {
//			t.Fatalf("GetRequest failed: %v", err)
//		}
//		expectedURL := "http://example.com/test"
//		if req.URL.String() != expectedURL {
//			t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
//		}
//	})
//
//}
