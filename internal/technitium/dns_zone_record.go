package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
)

func (c *Client) GetDnsZoneRecords(domain string, ctx context.Context) ([]DnsZoneRecord, error) {
	url := fmt.Sprintf("%s/api/zones/records/get?domain=%s", c.HostURL, domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, string(body))

	response := DnsZoneRecordsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Records, nil
}
