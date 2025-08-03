package technitium

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"strings"
)

func (c *Client) GetDnsZoneRecords(domain string, ctx context.Context) ([]DnsZoneRecord, error) {
	url := fmt.Sprintf("%s/api/zones/records/get?domain=%s", c.HostURL, domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("listZone", "true")
	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return nil, err
	}

	//tflog.Debug(ctx, string(body))

	response := DnsZoneRecordsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Records, nil
}

func (c *Client) GetDnsZoneRecord(domain string, recordType string, ctx context.Context) (DnsZoneRecord, error) {
	records, err := c.GetDnsZoneRecords(domain, ctx)

	if err != nil {
		return DnsZoneRecord{}, fmt.Errorf("failed to get DNS zone record: %w", err)
	}

	if len(records) == 0 {
		return DnsZoneRecord{}, fmt.Errorf("no DNS zone records found for domain: %s", domain)
	}

	for _, record := range records {
		if record.Type == recordType && record.Name == domain {
			tflog.Debug(ctx, fmt.Sprintf("Found DNS zone record: %+v", record))
			return record, nil
		}
	}

	return DnsZoneRecord{}, fmt.Errorf("no DNS zone record found for domain: %s with type: %s", domain, recordType)

}

func (c *Client) CreateDnsZoneRecord(r DnsZoneRecordCreate, ctx context.Context) error {

	req, err := c.GetRequest("/api/zones/records/add")
	if err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add("domain", r.Domain)
	params.Add("type", r.Type)
	params.Add("zone", r.Zone)
	params.Add("ttl", fmt.Sprintf("%d", r.TTL))
	params.Add("comments", r.Comments)
	params.Add("expiryTtl", fmt.Sprintf("%d", r.ExpiryTTL))
	params.Add("ipAddress", r.IPAddress)
	params.Add("ptr", r.Ptr)
	params.Add("createPtrZone", fmt.Sprintf("%t", r.CreatePtrZone))
	params.Add("updateSvcbHints", fmt.Sprintf("%t", r.UpdateSvcbHints))
	params.Add("nameServer", r.NameServer)
	params.Add("cname", r.Cname)
	params.Add("ptrName", r.PtrName)
	params.Add("exchange", r.Exchange)
	params.Add("preference", fmt.Sprintf("%d", r.Preference))
	params.Add("text", r.Text)
	params.Add("splitText", r.SplitText)
	params.Add("protocol", r.Protocol)
	params.Add("forwarder", r.Forwarder)
	params.Add("forwarderPriority", fmt.Sprintf("%d", r.ForwarderPriority))
	params.Add("dnssecValidation", fmt.Sprintf("%t", r.DnssecValidation))
	params.Add("proxyType", r.ProxyType)
	params.Add("proxyAddress", r.ProxyAddress)
	params.Add("proxyPort", fmt.Sprintf("%d", r.ProxyPort))
	params.Add("proxyUsername", r.ProxyUsername)
	params.Add("proxyPassword", r.ProxyPassword)
	params.Add("appName", r.AppName)
	params.Add("classPath", r.ClassPath)
	params.Add("recordData", r.RecordData)

	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	response := BaseResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {
		return fmt.Errorf("failed to create zone record: %s", response.ErrorMessage)
	}

	return nil

}

func (c *Client) UpdateDnsZoneRecord(r DnsZoneRecordUpdate, ctx context.Context) error {
	req, err := c.GetRequest("/api/zones/records/update")
	if err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add("domain", r.Domain)
	params.Add("newDomain", r.NewDomain)
	params.Add("zone", r.Zone)
	params.Add("type", r.Type)
	params.Add("disable", fmt.Sprintf("%t", r.Disable))
	params.Add("ttl", fmt.Sprintf("%d", r.TTL))
	params.Add("comments", r.Comments)
	params.Add("expiryTtl", fmt.Sprintf("%d", r.ExpiryTTL))
	params.Add("ipAddress", r.IPAddress)
	params.Add("newIpAddress", r.NewIPAddress)
	params.Add("ptr", r.Ptr)
	params.Add("createPtrZone", fmt.Sprintf("%t", r.CreatePtrZone))
	params.Add("updateSvcbHints", fmt.Sprintf("%t", r.UpdateSvcbHints))
	params.Add("nameServer", r.NameServer)
	params.Add("newNameServer", r.NewNameServer)
	params.Add("cname", r.Cname)
	params.Add("ptrName", r.PtrName)
	params.Add("newPtrName", r.NewPtrName)
	params.Add("exchange", r.Exchange)
	params.Add("newExchange", r.NewExchange)
	params.Add("preference", fmt.Sprintf("%d", r.Preference))
	params.Add("newPreference", fmt.Sprintf("%d", r.NewPreference))
	params.Add("text", r.Text)
	params.Add("newText", r.NewText)
	params.Add("splitText", r.SplitText)
	params.Add("newSplitText", r.NewSplitText)
	params.Add("protocol", r.Protocol)
	params.Add("forwarder", r.Forwarder)
	params.Add("newForwarder", r.NewForwarder)
	params.Add("forwarderPriority", fmt.Sprintf("%d", r.ForwarderPriority))
	params.Add("dnssecValidation", fmt.Sprintf("%t", r.DnssecValidation))
	params.Add("proxyType", r.ProxyType)
	params.Add("proxyAddress", r.ProxyAddress)
	params.Add("proxyPort", fmt.Sprintf("%d", r.ProxyPort))
	params.Add("proxyUsername", r.ProxyUsername)
	params.Add("proxyPassword", r.ProxyPassword)
	params.Add("appName", r.AppName)
	params.Add("classPath", r.ClassPath)
	params.Add("recordData", r.RecordData)

	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	response := BaseResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {
		return fmt.Errorf("failed to update zone record: %s", response.ErrorMessage)
	}

	return nil
}

func (c *Client) DeleteDnsZoneRecord(r DnsZoneRecordCreate, ctx context.Context) error {

	req, err := c.GetRequest("/api/zones/records/delete")
	if err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add("domain", r.Domain)
	params.Add("type", r.Type)
	params.Add("zone", r.Zone)
	params.Add("ipAddress", r.IPAddress)
	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	response := BaseResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {

		if strings.Contains(response.ErrorMessage, "Cannot delete record: no such record exists") {
			return nil
		}

		return fmt.Errorf("failed to delete zone record: %s", response.ErrorMessage)
	}

	return nil
}
