package serviceimpl

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"server/libs"
	"server/service"
	"sync"

	"github.com/gorilla/websocket"
)

type clientService struct {
	Clients map[string]*websocket.Conn
	token   string
	Mu      *sync.Mutex
}

func NewClientService(opt *libs.Options) service.ClientService {
	return &clientService{
		token:   opt.Token,
		Mu:      &sync.Mutex{},
		Clients: make(map[string]*websocket.Conn, 0),
	}
}

func (c clientService) BuildClient(input service.BuildClientInput) (string, error) {
	const buildStr = `GO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-X main.Port=%s -X main.ServerAddress=%s -X main.Token=%s -extldflags "-static"' -o ./%s main.go`

	buildCmd := fmt.Sprintf(buildStr, input.ServerPort, input.ServerAddress, c.token, input.Filename)
	log.Println("xxxxxx:", buildCmd)
	cmd := exec.Command("sh", "-c", buildCmd)
	cmd.Dir = "../client"

	outputErr, err := cmd.CombinedOutput()
	log.Println(outputErr)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("%w:%s", err, outputErr)
	}
	return input.Filename, nil
}

func (c clientService) AddConnection(clientID string, connection *websocket.Conn) error {
	c.Mu.Lock()
	c.Clients[clientID] = connection
	c.Mu.Unlock()
	return nil

}

func (c clientService) GetConnection(clientID string) (*websocket.Conn, bool) {
	c.Mu.Lock()
	conn, found := c.Clients[clientID]
	c.Mu.Unlock()
	return conn, found
}

func (c clientService) SendCommand(ctx context.Context, input service.SendCommandInput) ([]byte, error) {
	client, found := c.GetConnection(input.ClientID)
	if !found {
		return []byte("Connection NOT Found"), nil
	}
	log.Println("[!input command]: ", input.Command)
	err := client.WriteMessage(websocket.BinaryMessage, []byte(input.Command))
	if err != nil {
		log.Println(err)
		return []byte(err.Error()), err
	}

	_, readMessage, err := client.ReadMessage()
	if err != nil {
		log.Println(err)
		return []byte(err.Error()), err
	}

	return readMessage, nil
}
