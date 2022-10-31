package connect

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Command struct {
	ClientID  string `json:"client_id,omitempty"`
	Command   string `json:"command,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Response  []byte `json:"response,omitempty"`
	HasError  bool   `json:"has_error,omitempty"`
}

type CmdResponse struct {
	ClientID  string `json:"client_id,omitempty"`
	Command   string `json:"command,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Response  []byte `json:"response,omitempty"`
	Falco     bool
	HasError  bool `json:"has_error,omitempty"`
}

func (c *connector) reconnect() {
	c.connected = false
	for {
		conn, err := c.NewWebsocket()
		if err != nil {
			log.Println("[!] Error connecting on WS: ", err.Error())
			time.Sleep(time.Second * 10)
			continue
		}
		c.websockconn = conn
		c.connected = true
		log.Println(("[*] Successfully connected"))
		break
	}
}

func (c *connector) HandleCommand() {
	for {
		if !c.connected {
			c.reconnect()
			continue
		}

		// go c.SendFalco()

		_, message, err := c.websockconn.ReadMessage()
		if err != nil {
			log.Println("[!] Error reading from connection:", err)
			c.reconnect()
			continue
		}
		log.Println("[!message]: ", string(message))
		command := string(message)

		var response string
		var hasError bool
		switch command {
		case "falco restart":
			success, err := RestartFalco()
			if err != nil {
				hasError = true
				response = err.Error()
				continue
			}
			if success {
				response = "Restart Successed"
			} else {
				response = "Restart Fail"
			}
			break
		case "ifconfig":
			cmd := "ifconfig"
			resp, err := RunCommandWithErr(cmd)
			if err != nil {
				hasError = true
				resp = err.Error()
				continue
			}
			response = resp
			break

		}

		body, err := json.Marshal(&CmdResponse{
			ClientID: c.clientId,
			Response: []byte(response),
			HasError: hasError,
		})
		if err != nil {
			continue
		}
		log.Println("[!respons message]: ", string(body))
		err = c.websockconn.WriteMessage(websocket.BinaryMessage, body)
		if err != nil {
			continue
		}
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
	log.Println(timeout)
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
	log.Println("running command: ", cmd)
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

func RunCommandWithoutOutput(cmd string) error {
	command := []string{
		"bash",
		"-c",
		cmd,
	}
	log.Println("[Exec]: ", command)
	realCmd := exec.Command(command[0], command[1:]...)
	cmdReader, _ := realCmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()
	if err := realCmd.Start(); err != nil {
		return err
	}
	if err := realCmd.Wait(); err != nil {
		return err
	}
	log.Println("xxxxxxxxxxxxxx")
	return nil
}
