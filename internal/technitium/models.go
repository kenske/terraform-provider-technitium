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
	DnsSecStatus string `json:"dnssecStatus"`
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
	DnssecValidation           bool     `json:"dnssecValidation"`
}

type DnsZoneCreateResponse struct {
	Response struct {
		Domain string `json:"domain"`
	} `json:"response"`
	BaseResponse
}
