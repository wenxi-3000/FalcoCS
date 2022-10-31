package service

import (
	"context"

	"github.com/gorilla/websocket"
)

type CommandOutput struct {
	ClientID string `json:"client_id"`
	Response string `json:"response"`
	HasError bool   `json:"hasError"`
}

type BuildClientInput struct {
	ServerAddress, ServerPort, Filename string
}
type SendCommandInput struct {
	ClientID  string
	Command   string
	Parameter string
	Request   string
}

type ClientService interface {
	BuildClient(BuildClientInput) (string, error)
	AddConnection(clientID string, connection *websocket.Conn) error
	SendCommand(ctx context.Context, input SendCommandInput) ([]byte, error)
	GetConnection(clientID string) (*websocket.Conn, bool)
}
