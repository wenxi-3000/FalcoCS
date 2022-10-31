package entity

type GenerateClient struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	Filename string `json:"filename"`
}
