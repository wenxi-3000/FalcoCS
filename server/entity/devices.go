package entity

type Device struct {
	DBModel
	NodeIP     string
	Hostname   string
	MacAddress string
	ClientIPS  string
}
