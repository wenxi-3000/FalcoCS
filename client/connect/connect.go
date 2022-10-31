package connect

import (
	"client/device"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Connector interface {
	KeepConnection()
	HandleCommand()
}

type connector struct {
	timeout     time.Duration
	gerrors     chan error
	charset     string
	ipPort      string
	websockconn *websocket.Conn
	connected   bool
	token       string
	httpClient  *http.Client
	rootUri     string
	clientId    string
}

func NewConnector(ip string, port string, token string) Connector {
	c := connector{
		gerrors: make(chan error),
		timeout: 30 * time.Second,
		charset: "utf-8",
		ipPort:  ip + ":" + port,
		token:   token,
	}
	c.rootUri = "http://" + c.ipPort
	c.httpClient = NewHttpClient(10)
	c.clientId, _ = device.GetMacAddress()

	return &c
}

func NewHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * timeout,
	}
}

func (c *connector) KeepConnection() {
	sleepTime := 30 * time.Second

	for {
		if c.connected {
			time.Sleep(sleepTime)
		}

		err := c.health()
		if err != nil {
			log.Println("[!] Error connecting with server: " + err.Error())
			c.connected = false
			time.Sleep(sleepTime)
			continue
		}

		err = c.sender()
		if err != nil {
			log.Println("[!] Error connecting with server: " + err.Error())
			c.connected = false
			time.Sleep(sleepTime)
			continue
		}

		err = c.SendFalco()
		if err != nil {
			log.Println("[!] Error send falco: " + err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		c.connected = true
	}
}

//判断服务端是否能正常访问
func (c *connector) health() error {
	url := c.rootUri + "/health"

	log.Println(url)
	resp, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if resp.code != http.StatusOK {
		return fmt.Errorf(string(resp.body))
	}
	log.Println("[!]health check status OK: ", resp.code)
	return nil
}

//发送采集的客户端信息到服务端
func (c *connector) sender() error {
	device := device.NewDevice()
	url := c.rootUri + "/device"
	resp, err := c.NewRequest(http.MethodPost, url, device)
	if err != nil {
		return err
	}

	if resp.code != http.StatusOK {
		return fmt.Errorf("error with status code %d", resp.code)
	}
	return nil
}
