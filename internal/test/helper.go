package test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Scenario struct {
	ExpectedStatus int
	ExpectedBody   string
	ExpectedError  error // For simulating errors in DoFunc
}

// errorRoundTripper is a custom http.RoundTripper that always returns a configured error.
type errorRoundTripper struct {
	err error
}

// RoundTrip implements the http.RoundTripper interface.
func (rt *errorRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, rt.err
}

func GetTestClientComponents(scenario Scenario) (string, *http.Client, func()) {
	// If a client-side error is expected, mock the transport layer to return that error.
	if scenario.ExpectedError != nil {
		mockHTTPClient := &http.Client{
			Transport: &errorRoundTripper{err: scenario.ExpectedError},
		}
		// URL doesn't matter as the request won't be sent.
		// Return a no-op cleanup function as there's no server to close.
		return "http://mock-server.local", mockHTTPClient, func() {}
	}

	// Use httptest.Server to simulate a real API response.
	testServer := NewTestServer(scenario)

	// Return the server's URL, its client, and its Close function.
	return testServer.URL, testServer.Client(), testServer.Close
}

func NewTestServer(scenario Scenario) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(scenario.ExpectedStatus)
		fmt.Fprint(w, scenario.ExpectedBody)
	}))
}

func GetMockScenarioFromFile(t *testing.T, path string, statusCode int) Scenario {

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read mock file: %v", err)
	}

	scenario := Scenario{
		ExpectedStatus: statusCode,
		ExpectedBody:   string(data),
	}

	return scenario
}

func PrintPrettyDeepEqualError(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("Record mismatch:\nExpected: %s\nGot:      %s", spew.Sdump(expected), spew.Sdump(actual))
}
