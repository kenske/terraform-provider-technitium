package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
	"strings"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:5380"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

type StatusResponse struct {
	status string
}

func NewClient(host, token string, ctx context.Context) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != "" {
		c.HostURL = host
	}

	c.Token = token
	err := c.GetSessionInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &c, nil
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

	sr := StatusResponse{}
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
		tflog.Info(ctx, url)
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

func (c *Client) GetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, path), nil)

	if err != nil {
		return nil, err
	}

	return req, nil

}
