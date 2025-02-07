package technitium

type Scope struct {
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	StartingAddress  string `json:"starting_address"`
	EndingAddress    string `json:"ending_address"`
	SubnetMask       string `json:"subnet_mask"`
	NetworkAddress   string `json:"network_address"`
	BroadcastAddress string `json:"broadcast_address"`
	InterfaceAddress string `json:"interface_address"`
}
