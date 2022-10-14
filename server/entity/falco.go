package entity

type Falco struct {
	DBModel
	NodeIP string
	Falco  bool
}

type ParseHttpFalco struct {
	IPs   []string `json:"ips"`
	Falco bool
}
