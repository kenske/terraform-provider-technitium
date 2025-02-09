package technitium

type Scope struct {
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	StartingAddress  string `json:"startingAddress"`
	EndingAddress    string `json:"endingAddress"`
	SubnetMask       string `json:"subnetMask"`
	NetworkAddress   string `json:"networkAddress"`
	BroadcastAddress string `json:"broadcastAddress"`
	InterfaceAddress string `json:"interfaceAddress"`
}

type ScopesResponse struct {
	Response struct {
		Scopes []Scope
	} `json:"response"`
	Status string `json:"status"`
}
