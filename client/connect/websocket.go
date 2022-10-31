package connect

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func (c *connector) NewWebsocket() (*websocket.Conn, error) {
	host := c.ipPort

	scheme := "ws"

	u := url.URL{Scheme: scheme, Host: host, Path: "/client"}
	log.Println(u)
	header := http.Header{}
	header.Set("x-client", c.clientId)
	header.Set("Cookie", c.token)
	log.Println()
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	return conn, err
}

func (c *connector) FalcoWebsocket() (*websocket.Conn, error) {
	host := c.ipPort

	scheme := "ws"

	u := url.URL{Scheme: scheme, Host: host, Path: "/falco"}
	log.Println(u)
	header := http.Header{}
	header.Set("x-client", c.clientId)
	header.Set("Cookie", c.token)
	log.Println()
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	return conn, err
}
