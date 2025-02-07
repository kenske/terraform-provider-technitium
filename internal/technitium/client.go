package technitium

import (
	"encoding/json"
	"fmt"
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

func NewClient(host, token string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != "" {
		c.HostURL = host
	}

	c.Token = token
	err := c.GetSessionInfo()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) GetSessionInfo() error {
	if c.Token == "" {
		return fmt.Errorf("missing API token")
	}
	rb, err := json.Marshal(c)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/user/session/get?token=%s", c.HostURL, c.Token), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.URL.RawQuery = "?token=" + c.Token

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
