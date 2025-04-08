package technitium

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) CreateLease(s DhcpReservedLease, ctx context.Context) error {

	req, err := c.GetRequest("/api/dhcp/scopes/addReservedLease")

	if err != nil {
		return err
	}

	params := req.URL.Query()

	params.Add("name", s.Name)
	params.Add("hardwareAddress", s.HardwareAddress)
	params.Add("ipAddress", s.IpAddress)
	if s.HostName != "" {
		params.Add("hostName", s.HostName)
	}
	if s.Comments != "" {
		params.Add("comments", s.Comments)
	}

	req.URL.RawQuery = params.Encode()

	body, err := c.doRequest(req, ctx)
	if err != nil {
		return err
	}

	response := DhcpReservedLeaseResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {
		return fmt.Errorf("failed to create scope: %s", response.ErrorMessage)
	}

	return nil
}

func (c *Client) DeleteLease(name string, hardwareAddress string, ctx context.Context) error {

	req, err := c.GetRequest("/api/dhcp/scopes/removeReservedLease")

	if err != nil {
		return err
	}

	params := req.URL.Query()

	params.Add("name", name)
	params.Add("hardwareAddress", hardwareAddress)
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
		return fmt.Errorf("failed to delete lease: %s", response.ErrorMessage)
	}

	return nil
}
