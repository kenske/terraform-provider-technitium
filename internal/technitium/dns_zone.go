package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"strings"
)

func (c *Client) GetDnsZones(ctx context.Context) ([]DnsZoneList, error) {
	url := fmt.Sprintf("%s/api/zones/list", c.HostURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, string(body))

	response := DnsZonesResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Zones, nil
}

func (c *Client) GetDnsZone(name string, ctx context.Context) (DnsZone, error) {
	url := fmt.Sprintf("%s/api/zones/options/get?zone=%s", c.HostURL, name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DnsZone{}, err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return DnsZone{}, err
	}

	response := DnsZoneResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DnsZone{}, err
	}

	return response.Response, nil
}

func (c *Client) CreateDnsZone(z DnsZoneCreate, ctx context.Context) (DnsZone, error) {

	req, err := c.GetRequest("/api/zones/create")

	if err != nil {
		return DnsZone{}, err
	}

	params := req.URL.Query()

	params.Add("zone", z.Name)
	params.Add("type", z.Type)
	params.Add("catalog", z.Catalog)
	params.Add("useSoaSerialDateScheme", fmt.Sprintf("%v", z.UseSoaSerialDateScheme))
	params.Add("primaryNameServerAddresses", strings.Join(z.PrimaryNameServerAddresses, ","))
	params.Add("zoneTransferProtocol", z.ZoneTransferProtocol)
	params.Add("tsigKeyName", z.TsigKeyName)
	params.Add("protocol", z.Protocol)
	params.Add("forwarder", z.Forwarder)
	params.Add("dnssecValidation", fmt.Sprintf("%v", z.DnssecValidation))

	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return DnsZone{}, err
	}

	response := DnsZoneCreateResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DnsZone{}, err
	}

	if response.Status != "ok" {
		return DnsZone{}, fmt.Errorf("failed to create scope: %s", response.ErrorMessage)
	}

	return c.GetDnsZone(z.Name, ctx)

}
