package connect

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
