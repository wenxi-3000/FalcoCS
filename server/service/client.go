package service

type BuildClientInput struct {
	ServerAddress, ServerPort, Filename string
}

type ClientService interface {
	BuildClient(BuildClientInput) (string, error)
}
