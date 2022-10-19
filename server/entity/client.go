package entity

type GenerateClient struct {
	Address  string `form:"address"`
	Port     string `form:"port"`
	Filename string `form:"filename"`
}
