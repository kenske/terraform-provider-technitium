package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c *Client) GetScopes(ctx context.Context) ([]DhcpScopeList, error) {
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

	// Lease time parameters
	if s.LeaseTimeDays > 0 {
		params.Add("leaseTimeDays", fmt.Sprintf("%d", s.LeaseTimeDays))
	}
	if s.LeaseTimeHours > 0 {
		params.Add("leaseTimeHours", fmt.Sprintf("%d", s.LeaseTimeHours))
	}
	if s.LeaseTimeMinutes > 0 {
		params.Add("leaseTimeMinutes", fmt.Sprintf("%d", s.LeaseTimeMinutes))
	}

	// Offer delay and ping check
	if s.OfferDelayTime > 0 {
		params.Add("offerDelayTime", fmt.Sprintf("%d", s.OfferDelayTime))
	}
	params.Add("pingCheckEnabled", fmt.Sprintf("%v", s.PingCheckEnabled))
	if s.PingCheckTimeout > 0 {
		params.Add("pingCheckTimeout", fmt.Sprintf("%d", s.PingCheckTimeout))
	}
	if s.PingCheckRetries > 0 {
		params.Add("pingCheckRetries", fmt.Sprintf("%d", s.PingCheckRetries))
	}

	// Network configuration
	params.Add("routerAddress", s.RouterAddress)
	params.Add("useThisDnsServer", fmt.Sprintf("%v", s.UseThisDnsServer))
	params.Add("dnsServers", strings.Join(s.DnsServers, ","))
	params.Add("winsServers", strings.Join(s.WinsServers, ","))
	params.Add("ntpServers", strings.Join(s.NtpServers, ","))
	params.Add("ntpServerDomainNames", strings.Join(s.NtpServerDomainNames, ","))

	// Static routes
	if len(s.StaticRoutes) > 0 {
		routes := make([]string, 0)
		for _, route := range s.StaticRoutes {
			entry := fmt.Sprintf("%s|%s|%s", route.Destination, route.SubnetMask, route.Router)
			routes = append(routes, entry)
		}
		params.Add("staticRoutes", strings.Join(routes, "|"))
	}

	// Vendor info
	if len(s.VendorInfo) > 0 {
		vendors := make([]string, 0)
		for _, vendor := range s.VendorInfo {
			entry := fmt.Sprintf("%s|%s", vendor.Identifier, vendor.Information)
			vendors = append(vendors, entry)
		}
		params.Add("vendorInfo", strings.Join(vendors, "|"))
	}

	// Additional server options
	params.Add("capwapAcIpAddresses", strings.Join(s.CAPWAPAcIpAddresses, ","))
	params.Add("tftpServerAddresses", strings.Join(s.TftpServerAddresses, ","))

	// Generic options
	if len(s.GenericOptions) > 0 {
		options := make([]string, 0)
		for _, option := range s.GenericOptions {
			entry := fmt.Sprintf("%d|%s", option.Code, option.Value)
			options = append(options, entry)
		}
		params.Add("genericOptions", strings.Join(options, "|"))
	}

	// Domain configuration
	params.Add("domainName", s.DomainName)
	params.Add("domainSearchList", strings.Join(s.DomainSearchList, ","))

	// Boot options
	params.Add("bootFileName", s.BootFileName)
	params.Add("nextServerAddress", s.NextServerAddress)
	params.Add("serverHostName", s.ServerHostName)
	params.Add("serverAddress", s.ServerAddress)

	// Interface binding
	params.Add("interfaceAddress", s.InterfaceAddress)
	if s.InterfaceIndex > 0 {
		params.Add("interfaceIndex", fmt.Sprintf("%d", s.InterfaceIndex))
	}

	// Reserved leases and access control
	params.Add("reservedLeases", strings.Join(s.ReservedLeases, ","))
	params.Add("allowOnlyReservedLeases", fmt.Sprintf("%v", s.AllowOnlyReservedLeases))
	params.Add("blockLocallyAdministeredMacAddresses", fmt.Sprintf("%v", s.BlockLocallyAdministeredMacAddresses))
	params.Add("ignoreClientIdentifierOption", fmt.Sprintf("%v", s.IgnoreClientIdentifierOption))

	if oldName != "" {
		tflog.Debug(ctx, fmt.Sprintf("Renaming scope from %s to %s", oldName, s.Name))
		params.Set("name", oldName)
		params.Add("newName", s.Name)
	} else {
		tflog.Debug(ctx, fmt.Sprintf("Creating new scope %s", s.Name))
	}

	// Exclusions
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
