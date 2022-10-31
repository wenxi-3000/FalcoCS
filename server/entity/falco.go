package entity

type Falco struct {
	DBModel
	ClientID string `json:"client_id"`
	NodeIp   string
	Falco    bool `json:"falco"`
}

type ParseHttpFalco struct {
	IPs   []string `json:"ips"`
	Falco bool
}
