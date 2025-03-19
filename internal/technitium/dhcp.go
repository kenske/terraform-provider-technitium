package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"strings"
)

func (c *Client) GetScopes(ctx context.Context) ([]DhcpListScope, error) {
	url := fmt.Sprintf("%s/api/dhcp/scopes/list", c.HostURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return nil, err
	}

	response := DhcpScopesResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Scopes, nil
}

func (c *Client) GetScope(name string, ctx context.Context) (DhcpScope, error) {
	url := fmt.Sprintf("%s/api/dhcp/scopes/get?name=%s", c.HostURL, name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DhcpScope{}, err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return DhcpScope{}, err
	}

	response := DhcpScopeResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DhcpScope{}, err
	}

	return response.Response, nil
}

func (c *Client) SetScope(s DhcpScope, oldName string, ctx context.Context) (DhcpScope, error) {

	req, err := c.GetRequest("/api/dhcp/scopes/set")

	if err != nil {
		return DhcpScope{}, err
	}

	params := req.URL.Query()

	params.Add("name", s.Name)
	params.Add("startingAddress", s.StartingAddress)
	params.Add("endingAddress", s.EndingAddress)
	params.Add("subnetMask", s.SubnetMask)
	params.Add("routerAddress", s.RouterAddress)
	params.Add("useThisDnsServer", fmt.Sprintf("%v", s.UseThisDnsServer))
	params.Add("dnsServers", strings.Join(s.DnsServers, ","))
	params.Add("domainName", s.DomainName)

	if oldName != "" {
		tflog.Debug(ctx, fmt.Sprintf("Renaming scope from %s to %s", oldName, s.Name))
		params.Set("name", oldName)
		params.Add("newName", s.Name)
	} else {
		tflog.Debug(ctx, fmt.Sprintf("Creating new scope %s", s.Name))
	}

	exclusions := make([]string, 0)
	if s.Exclusions != nil {
		for _, exclusion := range s.Exclusions {
			entry := fmt.Sprintf("%s|%s", exclusion.StartingAddress, exclusion.EndingAddress)
			exclusions = append(exclusions, entry)
		}

		params.Add("exclusions", strings.Join(exclusions, "|"))
	}

	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return DhcpScope{}, err
	}

	response := DhcpScopeResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DhcpScope{}, err
	}

	if response.Status != "ok" {
		return DhcpScope{}, fmt.Errorf("failed to create scope: %s", response.ErrorMessage)
	}

	return c.GetScope(s.Name, ctx)
}

func (c *Client) DeleteScope(name string, ctx context.Context) error {

	req, err := c.GetRequest("/api/dhcp/scopes/delete")

	if err != nil {
		return err
	}

	params := req.URL.Query()

	params.Add("name", name)
	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	response := DhcpScopeResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}
