package technitium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// tokenQueryParamPattern matches a "token" query string parameter so its
// value can be redacted before a URL is logged or surfaced in an error.
var tokenQueryParamPattern = regexp.MustCompile(`token=[^&]+`)

// redactTokenQueryParam replaces the value of a "token" query string
// parameter with a fixed placeholder so a live API token never ends up in
// logs or error messages (LegacyTokenAuth mode carries the token in the URL).
func redactTokenQueryParam(s string) string {
	return tokenQueryParamPattern.ReplaceAllString(s, "token=hidden")
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HostURL    string
	HTTPClient HttpClient
	Token      string

	// LegacyTokenAuth sends the API token as a "token" URL query parameter
	// instead of an "Authorization: Bearer" header. Technitium DNS Server
	// versions prior to 15.0 only support the query parameter form. Leaving
	// this false (the default) keeps the token out of request URLs, which
	// commonly end up in reverse proxy access logs.
	LegacyTokenAuth bool
}

func NewClient(host, token string, legacyTokenAuth bool, ctx context.Context) (*Client, error) {
	c := Client{
		HTTPClient:      &http.Client{Timeout: 10 * time.Second},
		HostURL:         host,
		LegacyTokenAuth: legacyTokenAuth,
	}

	c.Token = token
	err := c.GetSessionInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetToken(host string, username string, password string) (string, error) {

	if username == "" || password == "" {
		return "", fmt.Errorf("username and password must be provided")
	}

	// Send credentials as a POST form body rather than a URL query string.
	// The Technitium API accepts both GET and POST for /api/user/login, and
	// only the POST form avoids putting the username and password in the
	// request URL, where they would otherwise be captured in cleartext by
	// reverse proxy access logs, browser history, and *url.Error messages
	// from transport failures.
	form := url.Values{}
	form.Set("user", username)
	form.Set("pass", password)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/user/login", host), strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	token, ok := result["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

func (c *Client) GetSessionInfo(ctx context.Context) error {
	if c.Token == "" {
		return fmt.Errorf("missing API token")
	}

	rb, err := json.Marshal(c.Token)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/user/session/get", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	sr := BaseResponse{}
	err = json.Unmarshal(body, &sr)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doRequest(req *http.Request, ctx context.Context) ([]byte, error) {

	if c.LegacyTokenAuth {
		// append token to url parameters (pre-15.0 Technitium DNS Server)
		query := req.URL.Query()
		query.Add("token", c.Token)
		req.URL.RawQuery = query.Encode()
	} else {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	if ctx != nil {
		// Hide token in the URL for logging
		tflog.Info(ctx, redactTokenQueryParam(req.URL.String()))
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, redactTokenFromError(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// redactTokenFromError strips a live API token out of a transport error
// before it is returned to the caller. Go's *http.Client.Do wraps every
// transport failure (DNS failure, connection refused, TLS failure, timeout)
// in a *url.Error whose Error() string embeds the full request URL,
// including its query string. In LegacyTokenAuth mode that query string
// carries the "token" parameter, so an unredacted transport error would leak
// the live token into whatever surfaces err.Error() -- e.g. a Terraform
// diagnostic message shown in a user's terminal or CI log.
func redactTokenFromError(err error) error {
	var uerr *url.Error
	if errors.As(err, &uerr) {
		redacted := *uerr
		redacted.URL = redactTokenQueryParam(uerr.URL)
		return &redacted
	}

	return err
}

// GetRequest TODO: change this function to accept a map with GET params
func (c *Client) GetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, path), nil)

	if err != nil {
		return nil, err
	}

	return req, nil

}
