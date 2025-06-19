package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type dhcpScopeList struct {
	Name             types.String `tfsdk:"name"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	StartingAddress  types.String `tfsdk:"starting_address"`
	EndingAddress    types.String `tfsdk:"ending_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	NetworkAddress   types.String `tfsdk:"network_address"`
	BroadcastAddress types.String `tfsdk:"broadcast_address"`
}

type dhcpScope struct {
	Name             types.String   `tfsdk:"name"`
	StartingAddress  types.String   `tfsdk:"starting_address"`
	EndingAddress    types.String   `tfsdk:"ending_address"`
	SubnetMask       types.String   `tfsdk:"subnet_mask"`
	RouterAddress    types.String   `tfsdk:"router_address"`
	UseThisDnsServer types.Bool     `tfsdk:"use_this_dns_server"`
	DnsServers       []types.String `tfsdk:"dns_servers"`
	DomainName       types.String   `tfsdk:"domain_name"`
	Exclusions       []Exclusion    `tfsdk:"exclusions"`
}

type Exclusion struct {
	StartingAddress types.String `tfsdk:"starting_address"`
	EndingAddress   types.String `tfsdk:"ending_address"`
}

type dhcpReservedLease struct {
	Name            types.String `tfsdk:"name"`
	HardwareAddress types.String `tfsdk:"hardware_address"`
	IpAddress       types.String `tfsdk:"ip_address"`
	HostName        types.String `tfsdk:"host_name"`
	Comments        types.String `tfsdk:"comments"`
}

type dnsZone struct {
	Name    types.String `tfsdk:"name"`
	Type    types.String `tfsdk:"type"`
	Catalog types.String `tfsdk:"catalog"`
}

type dnsZoneGet struct {
	dnsZone
	Disabled                 types.Bool     `tfsdk:"disabled"`
	DnssecStatus             types.String   `tfsdk:"dnssec_status"`
	NotifyFailed             types.Bool     `tfsdk:"notify_failed"`
	NotifyFailedFor          []types.String `tfsdk:"notify_failed_for"`
	QueryAccess              types.String   `tfsdk:"query_access"`
	QueryAccessNetworkAcl    []types.String `tfsdk:"query_access_network_acl"`
	ZoneTransfer             types.String   `tfsdk:"zone_transfer"`
	ZoneTransferNetworkAcl   []types.String `tfsdk:"zone_transfer_network_acl"`
	ZoneTransferTsigKeyNames []types.String `tfsdk:"zone_transfer_tsig_key_names"`
	Notify                   types.String   `tfsdk:"notify"`
	NotifyNameServers        []types.String `tfsdk:"notify_name_servers"`
	Update                   types.String   `tfsdk:"update"`
	UpdateNetworkAcl         []types.String `tfsdk:"update_network_acl"`
}

type dnsZoneList struct {
	dnsZone
	Disabled     types.Bool   `tfsdk:"disabled"`
	DnssecStatus types.String `tfsdk:"dnssec_status"`
	SoaSerial    types.Int32  `tfsdk:"soa_serial"`
	Expiry       types.String `tfsdk:"expiry"`
	IsExpired    types.Bool   `tfsdk:"is_expired"`
	LastModified types.String `tfsdk:"last_modified"`
	Internal     types.Bool   `tfsdk:"internal"`
}

type dnsZoneCreate struct {
	dnsZone
	Forwarder                  types.String   `tfsdk:"forwarder"`
	InitializeForwarder        types.Bool     `tfsdk:"initialize_forwarder"`
	UseSoaSerialDateScheme     types.Bool     `tfsdk:"use_soa_serial_date_scheme"`
	PrimaryNameServerAddresses []types.String `tfsdk:"primary_name_server_addresses"`
	ZoneTransferProtocol       types.String   `tfsdk:"zone_transfer_protocol"`
	TsigKeyName                types.String   `tfsdk:"tsig_key_name"`
	Protocol                   types.String   `tfsdk:"protocol"`
	DnssecValidation           types.Bool     `tfsdk:"dnssec_validation"`
}

type dnsZoneRecords struct {
	Domain  types.String    `tfsdk:"domain"`
	Records []dnsZoneRecord `tfsdk:"records"`
}

type dnsZoneRecord struct {
	Name         types.String       `tfsdk:"name"`
	Type         types.String       `tfsdk:"type"`
	TTL          types.Int64        `tfsdk:"ttl"`
	Disabled     types.Bool         `tfsdk:"disabled"`
	DnssecStatus types.String       `tfsdk:"dnssec_status"`
	LastUsedOn   types.String       `tfsdk:"last_used_on"`
	LastModified types.String       `tfsdk:"last_modified"`
	ExpiryTTL    types.Int64        `tfsdk:"expiry_ttl"`
	RecordData   *dnsZoneRecordData `tfsdk:"record_data"`
}

type dnsZoneRecordData struct {
	PrimaryNameServer   types.String `tfsdk:"primary_name_server"`
	ResponsiblePerson   types.String `tfsdk:"responsible_person"`
	Serial              types.Int64  `tfsdk:"serial"`
	Refresh             types.Int64  `tfsdk:"refresh"`
	Retry               types.Int64  `tfsdk:"retry"`
	Expire              types.Int64  `tfsdk:"expire"`
	Minimum             types.Int64  `tfsdk:"minimum"`
	UseSerialDateScheme types.Bool   `tfsdk:"use_serial_date_scheme"`
	Protocol            types.String `tfsdk:"protocol"`
	Forwarder           types.String `tfsdk:"forwarder"`
	Priority            types.Int64  `tfsdk:"priority"`
	DnssecValidation    types.Bool   `tfsdk:"dnssec_validation"`
	ProxyType           types.String `tfsdk:"proxy_type"`
	IpAddress           types.String `tfsdk:"ip_address"`
	Cname               types.String `tfsdk:"cname"`
	NameServer          types.String `tfsdk:"name_server"`
}

type dnsZoneRecordCreate struct {
	Domain            types.String `tfsdk:"domain"`
	Zone              types.String `tfsdk:"zone"`
	Type              types.String `tfsdk:"type"`
	TTL               types.Int64  `tfsdk:"ttl"`
	Comments          types.String `tfsdk:"comments"`
	ExpiryTTL         types.Int64  `tfsdk:"expiry_ttl"`
	IPAddress         types.String `tfsdk:"ip_address"`
	Ptr               types.String `tfsdk:"ptr"`
	CreatePtrZone     types.Bool   `tfsdk:"create_ptr_zone"`
	UpdateSvcbHints   types.Bool   `tfsdk:"update_svcb_hints"`
	NameServer        types.String `tfsdk:"name_server"`
	Cname             types.String `tfsdk:"cname"`
	PtrName           types.String `tfsdk:"ptr_name"`
	Exchange          types.String `tfsdk:"exchange"`
	Preference        types.Int64  `tfsdk:"preference"`
	Text              types.String `tfsdk:"text"`
	SplitText         types.String `tfsdk:"split_text"`
	Protocol          types.String `tfsdk:"protocol"`
	Forwarder         types.String `tfsdk:"forwarder"`
	ForwarderPriority types.Int64  `tfsdk:"forwarder_priority"`
	DnssecValidation  types.Bool   `tfsdk:"dnssec_validation"`
	ProxyType         types.String `tfsdk:"proxy_type"`
	ProxyAddress      types.String `tfsdk:"proxy_address"`
	ProxyPort         types.Int64  `tfsdk:"proxy_port"`
	ProxyUsername     types.String `tfsdk:"proxy_username"`
	ProxyPassword     types.String `tfsdk:"proxy_password"`
	AppName           types.String `tfsdk:"app_name"`
	ClassPath         types.String `tfsdk:"class_path"`
	RecordData        types.String `tfsdk:"record_data"`
}
