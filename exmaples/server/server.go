package server

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var conn net.Conn

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	//下发命令
	_, err := io.Copy(os.Stdin, conn)
	if err != nil {
		fmt.Println(err)
	}
	// reader := bufio.NewReader(os.Stdin)
	// data, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = conn.Write([]byte(data))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//读取信息
	// var buf [128]byte
	// n, err := conn.Read(buf[:])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(n)

	//
}

func Listen() {
	//监听
	var err error
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()
	buf := make([]byte, 4096)
	for {
		time.Sleep(time.Second * 1)
		conn, err = listener.Accept()
		if err != nil {
			log.Println(err)
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		filename := string(buf[:n])
		log.Println("filename:", filename)
	}

}

func Server() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//监听
	go Listen()

	//下发命令
	for {
		time.Sleep(time.Second * 1)
		if conn != nil {
			log.Println(conn.RemoteAddr())
			//cmd := "id"
			cmd := base64.URLEncoding.EncodeToString([]byte("id"))
			log.Println(cmd)
			conn.Write([]byte(cmd + "\n"))
		}

	}
	// for {
	// 	conn, err = server.Accept()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	message, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	decode, _ := base64.StdEncoding.DecodeString(message)
	// 	fmt.Println(decode)
	// }

	//开启监听
	// server, err := net.Listen("tcp", "0.0.0.0:8081")
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer server.Close()
	// fmt.Println("服务端监听已启动")

	// for {
	// 	//接收数据并启一个goroutine进行处理
	// 	conn, err := server.Accept()
	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	go HandleConnection(conn)
	// }

	// for {
	// 	command := ReadLine()
	// 	cmd := base64.URLEncoding.EncodeToString([]byte(command))
	// 	conn.Write([]byte(cmd + "\n"))
	// }
}

// ReadLine 函数等待命令行输入,返回字符串
func ReadLine() string {
	buf := bufio.NewReader(os.Stdin)
	lin, _, err := buf.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(lin)
}
