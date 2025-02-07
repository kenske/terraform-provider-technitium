package technitium

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetScopes() ([]Scope, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dhcp/scopes/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	scopes := []Scope{}
	err = json.Unmarshal(body, &scopes)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}
