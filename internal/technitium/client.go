package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HostURL    string
	HTTPClient HttpClient
	Token      string
}

func NewClient(host, token string, ctx context.Context) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    host,
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

	url := fmt.Sprintf("%s/api/user/login?user=%s&pass=%s", host, username, password)
	resp, err := http.Get(url)
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/user/session/get?token=%s", c.HostURL, c.Token), strings.NewReader(string(rb)))
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

	// append token to url parameters
	query := req.URL.Query()
	query.Add("token", c.Token)
	req.URL.RawQuery = query.Encode()

	if ctx != nil {
		url := req.URL.String()
		// Hide token in the URL for logging
		re := regexp.MustCompile(`token=[^&]+`)
		censoredURL := re.ReplaceAllString(url, "token=hidden")
		tflog.Info(ctx, censoredURL)
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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

// GetRequest TODO: change this function to accept a map with GET params
func (c *Client) GetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, path), nil)

	if err != nil {
		return nil, err
	}

	return req, nil

}
