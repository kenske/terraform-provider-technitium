package technitium

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetScopes() ([]Scope, error) {
	url := fmt.Sprintf("%s/api/dhcp/scopes/list", c.HostURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := ScopesResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Scopes, nil
}
