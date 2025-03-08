package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Client) CreateScope(s DhcpScope, ctx context.Context) (DhcpScope, error) {

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
	//params.Add("interfaceAddress", s.InterfaceAddress)
	//params.Add("ntpServers", fmt.Sprintf("%v", s.NtpServers))
	//params.Add("staticRoutes", fmt.Sprintf("%v", s.StaticRoutes))
	//params.Add("vendorInfo", fmt.Sprintf("%v", s.VendorInfo))
	//params.Add("capwapAcIpAddresses", fmt.Sprintf("%v", s.CapwapAcIpAddresses))
	//params.Add("tftpServerAddresses", fmt.Sprintf("%v", s.TftpServerAddresses))
	//params.Add("genericOptions", fmt.Sprintf("%v", s.GenericOptions))
	//params.Add("exclusions", fmt.Sprintf("%v", s.Exclusions))
	//params.Add("reservedLeases", fmt.Sprintf("%v", s.ReservedLeases))
	//params.Add("allowOnlyReservedLeases", fmt.Sprintf("%t", s.AllowOnlyReservedLeases))
	//params.Add("blockLocallyAdministeredMacAddresses", fmt.Sprintf("%t", s.BlockLocallyAdministeredMac))
	//params.Add("ignoreClientIdentifierOption", fmt.Sprintf("%t", s.IgnoreClientIdentifier))

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

	return response.Response, nil
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
