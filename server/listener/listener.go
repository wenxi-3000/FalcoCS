package listener

import (
	"bufio"
	"encoding/base64"
	"log"
	"net"
	"strings"
	"time"
)

type Listener struct {
	Connlist    map[string]net.Conn
	IPPort      string
	SendChan    chan string
	ReceiveChan chan string
}

func Newlistener(ip string, port string) *Listener {
	l := Listener{
		Connlist:    make(map[string]net.Conn),
		IPPort:      ip + ":" + port,
		SendChan:    make(chan string),
		ReceiveChan: make(chan string),
	}
	return &l
}

func ReadMessage(conn net.Conn) (result string) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	decoded, _ := base64.StdEncoding.DecodeString(message)
	decMessage := string(decoded)
	if message != "" {
		// log.Println(decMessage)
		return decMessage
	}
	return ""
}

// 等待Socket 客户端连接
func (listener *Listener) handleConnWait() {
	l, err := net.Listen("tcp", listener.IPPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// connlistx := make(map[string]net.Conn)
	for {
		time.Sleep(1 * time.Second)
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
		//connlistx[ip] = conn

		listener.Connlist[ip] = conn
		// connlistx[] = conn
		// r.connlist <- connlistx

		// go readMessage(conn)
		//message, err := bufio.NewReader(conn).ReadString('\n')
		// message, err := bufio.NewReader(conn).ReadString('\n')
		// decoded, _ := base64.StdEncoding.DecodeString(message)
		// go connection(conn)
	}
}

func InitListener(listener *Listener) {
	listener.handleConnWait()
}

func RunComand(command string, conn net.Conn) {
	cmd := base64.URLEncoding.EncodeToString([]byte(command))
	conn.Write([]byte(cmd + "\n"))
}

// for _, conn := range listener.connlist {
// 	if conn[ip] != "" {

// 	}
// 	_conn := conn[ip]
// 	log.Println(_conn)
// 	_cmd := base64.URLEncoding.EncodeToString([]byte(command))
// 	log.Println(_cmd)
// 	_conn.Write([]byte(_cmd + "\n"))

// }
