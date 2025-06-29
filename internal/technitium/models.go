package technitium

type BaseResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type DhcpScopeList struct {
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	StartingAddress  string `json:"startingAddress"`
	EndingAddress    string `json:"endingAddress"`
	SubnetMask       string `json:"subnetMask"`
	NetworkAddress   string `json:"networkAddress"`
	BroadcastAddress string `json:"broadcastAddress"`
}

type DhcpScopesResponse struct {
	Response struct {
		Scopes []DhcpScopeList
	} `json:"response"`
	BaseResponse
}

type DhcpScope struct {
	Name             string      `json:"name"`
	StartingAddress  string      `json:"startingAddress"`
	EndingAddress    string      `json:"endingAddress"`
	SubnetMask       string      `json:"subnetMask"`
	RouterAddress    string      `json:"routerAddress,omitempty"`
	UseThisDnsServer bool        `json:"useThisDnsServer,omitempty"`
	DnsServers       []string    `json:"dnsServers,omitempty"`
	DomainName       string      `json:"DomainName,omitempty"`
	Exclusions       []Exclusion `json:"exclusions"`
}

type Exclusion struct {
	StartingAddress string `json:"startingAddress"`
	EndingAddress   string `json:"endingAddress"`
}

type DhcpScopeResponse struct {
	Response DhcpScope `json:"response"`
	BaseResponse
}

type DhcpReservedLease struct {
	Name            string `json:"name"`
	HardwareAddress string `json:"hardwareAddress"`
	IpAddress       string `json:"ipAddress"`
	HostName        string `json:"hostName,omitempty"`
	Comments        string `json:"comments,omitempty"`
}

type DhcpReservedLeaseResponse struct {
	Response DhcpReservedLease `json:"response"`
	BaseResponse
}

type DnsZoneList struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Disabled     bool   `json:"disabled"`
	DnssecStatus string `json:"dnssecStatus"`
	SoaSerial    int    `json:"soaSerial"`
	Expiry       string `json:"expiry,omitempty"`
	IsExpired    bool   `json:"isExpired,omitempty"`
	LastModified string `json:"lastModified"`
	Internal     bool   `json:"internal"`
	Catalog      string `json:"catalog"`
}

type DnsZonesResponse struct {
	Response struct {
		PageNumber int           `json:"pageNumber"`
		TotalPages int           `json:"totalPages"`
		TotalZones int           `json:"totalZones"`
		Zones      []DnsZoneList `json:"zones"`
	} `json:"response"`
	BaseResponse
}

type DnsZone struct {
	Name                     string   `json:"name"`
	Type                     string   `json:"type"`
	Disabled                 bool     `json:"disabled"`
	DnsSecStatus             string   `json:"dnsSecStatus"`
	Catalog                  string   `json:"catalog"`
	NotifyFailed             bool     `json:"notifyFailed"`
	NotifyFailedFor          []string `json:"notifyFailedFor"`
	QueryAccess              string   `json:"queryAccess"`
	QueryAccessNetworkAcl    []string `json:"queryAccessNetworkACL"`
	ZoneTransfer             string   `json:"zoneTransfer"`
	ZoneTransferNetworkAcl   []string `json:"zoneTransferNetworkACL"`
	ZoneTransferTsigKeyNames []string `json:"zoneTransferTsigKeyNames"`
	Notify                   string   `json:"notify"`
	NotifyNameServers        []string `json:"notifyNameServers"`
	Update                   string   `json:"update"`
	UpdateNetworkAcl         []string `json:"updateNetworkACL"`
}

type DnsZoneResponse struct {
	Response DnsZone `json:"response"`
	BaseResponse
}

type DnsZoneCreate struct {
	Name                       string   `json:"name"`
	Type                       string   `json:"type"`
	Catalog                    string   `json:"catalog"`
	UseSoaSerialDateScheme     bool     `json:"useSoaSerialDateScheme"`
	PrimaryNameServerAddresses []string `json:"primaryNameServerAddresses"`
	ZoneTransferProtocol       string   `json:"zoneTransferProtocol"`
	TsigKeyName                string   `json:"tsigKeyName"`
	Protocol                   string   `json:"protocol"`
	Forwarder                  string   `json:"forwarder"`
	InitializeForwarder        bool     `json:"initializeForwarder"`
	DnssecValidation           bool     `json:"dnssecValidation"`
}

type DnsZoneCreateResponse struct {
	Response struct {
		Domain string `json:"domain"`
	} `json:"response"`
	BaseResponse
}

type DnsZoneRecordsResponse struct {
	Response struct {
		Records []DnsZoneRecord `json:"records"`
	} `json:"response"`
	BaseResponse
}

type DnsZoneRecord struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	TTL          int64             `json:"ttl"`
	Disabled     bool              `json:"disabled"`
	DnsSecStatus string            `json:"dnsSecStatus"`
	LastUsedOn   string            `json:"lastUsedOn"`
	LastModified string            `json:"lastModified"`
	ExpiryTTL    int64             `json:"expiryTTL"`
	RecordData   DnsZoneRecordData `json:"rData"`
}

type DnsZoneRecordData struct {
	PrimaryNameServer   string `json:"primaryNameServer,omitempty"`
	ResponsiblePerson   string `json:"responsiblePerson,omitempty"`
	Serial              int64  `json:"serial,omitempty"`
	Refresh             int64  `json:"refresh"`
	Retry               int64  `json:"retry"`
	Expire              int64  `json:"expire"`
	Minimum             int64  `json:"minimum"`
	UseSerialDateScheme bool   `json:"useSerialDateScheme"`
	Protocol            string `json:"protocol"`
	Forwarder           string `json:"forwarder"`
	Priority            int64  `json:"priority"`
	DnssecValidation    bool   `json:"dnssecValidation"`
	ProxyType           string `json:"proxyType"`
	IpAddress           string `json:"ipAddress"`
	NameServer          string `json:"nameServer"`
	Cname               string `json:"cname"`
}

type DnsZoneRecordCreate struct {
	Domain            string `json:"domain"`
	Type              string `json:"type"`
	Zone              string `json:"zone,omitempty"`
	TTL               int64  `json:"ttl,omitempty"`
	Comments          string `json:"comments,omitempty"`
	ExpiryTTL         int64  `json:"expiryTtl,omitempty"`
	IPAddress         string `json:"ipAddress,omitempty"`
	Ptr               string `json:"ptr,omitempty"`
	CreatePtrZone     bool   `json:"createPtrZone,omitempty"`
	UpdateSvcbHints   bool   `json:"updateSvcbHints,omitempty"`
	NameServer        string `json:"nameServer,omitempty"`
	Cname             string `json:"cname,omitempty"`
	PtrName           string `json:"ptrName,omitempty"`
	Exchange          string `json:"exchange,omitempty"`
	Preference        int64  `json:"preference,omitempty"`
	Text              string `json:"text,omitempty"`
	SplitText         string `json:"splitText,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Forwarder         string `json:"forwarder,omitempty"`
	ForwarderPriority int64  `json:"forwarderPriority,omitempty"`
	DnssecValidation  bool   `json:"dnssecValidation,omitempty"`
	ProxyType         string `json:"proxyType,omitempty"`
	ProxyAddress      string `json:"proxyAddress,omitempty"`
	ProxyPort         int64  `json:"proxyPort,omitempty"`
	ProxyUsername     string `json:"proxyUsername,omitempty"`
	ProxyPassword     string `json:"proxyPassword,omitempty"`
	AppName           string `json:"appName,omitempty"`
	ClassPath         string `json:"classPath,omitempty"`
	RecordData        string `json:"recordData,omitempty"`
}

type DnsZoneRecordUpdate struct {
	DnsZoneRecordCreate
	NewDomain     string `json:"newDomain,omitempty"`
	NewIPAddress  string `json:"newIPAddress,omitempty"`
	NewNameServer string `json:"newNameServer,omitempty"`
	NewPtrName    string `json:"newPtrName,omitempty"`
	NewExchange   string `json:"newExchange,omitempty"`
	NewPreference int64  `json:"newPreference,omitempty"`
	NewText       string `json:"newText,omitempty"`
	NewSplitText  string `json:"newSplitText,omitempty"`
	NewForwarder  string `json:"newForwarder,omitempty"`
}
