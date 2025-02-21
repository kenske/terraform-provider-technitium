package technitium

type ListDhcpScope struct {
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	StartingAddress  string `json:"startingAddress"`
	EndingAddress    string `json:"endingAddress"`
	SubnetMask       string `json:"subnetMask"`
	NetworkAddress   string `json:"networkAddress"`
	BroadcastAddress string `json:"broadcastAddress"`
	InterfaceAddress string `json:"interfaceAddress"`
}

type ListScopesResponse struct {
	Response struct {
		Scopes []ListDhcpScope
	} `json:"response"`
	Status string `json:"status"`
}

type DhcpScope struct {
	ListDhcpScope
	NtpServers   []string `json:"ntpServers"`
	StaticRoutes []struct {
		Destination string `json:"destination"`
		SubnetMask  string `json:"subnetMask"`
		Router      string `json:"router"`
	} `json:"staticRoutes"`
	VendorInfo []struct {
		Identifier  string `json:"identifier"`
		Information string `json:"information"`
	} `json:"vendorInfo"`
	CapwapAcIpAddresses []string `json:"capwapAcIpAddresses"`
	TftpServerAddresses []string `json:"tftpServerAddresses"`
	GenericOptions      []struct {
		Code  int32  `json:"code"`
		Value string `json:"value"`
	} `json:"genericOptions"`
	Exclusions []struct {
		StartingAddress string `json:"startingAddress"`
		EndingAddress   string `json:"endingAddress"`
	} `json:"exclusions"`
	ReservedLeases []struct {
		HostName        string `json:"hostName,omitempty"`
		HardwareAddress string `json:"hardwareAddress"`
		Address         string `json:"address"`
		Comments        string `json:"comments,omitempty"`
	} `json:"reservedLeases"`
	AllowOnlyReservedLeases     bool `json:"allowOnlyReservedLeases"`
	BlockLocallyAdministeredMac bool `json:"blockLocallyAdministeredMacAddresses"`
	IgnoreClientIdentifier      bool `json:"ignoreClientIdentifierOption"`
}

type DhcpScopeResponse struct {
	Response DhcpScope `json:"response"`
	Status   string    `json:"status"`
}
