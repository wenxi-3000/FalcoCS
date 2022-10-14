package falco

import (
	"client/device"
	"encoding/json"
	"log"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type Falco struct {
	IPs   []string `json:"ips"`
	Falco bool     `json:"falco"`
}

func NewFalco() []byte {
	var falco Falco

	falco.IPs = device.GetLaddrIP()
	falco.Falco = GetFalcoProcess()

	result, err := json.Marshal(falco)
	if err != nil {
		log.Println(err)
	}
	return result
}

// func InfoSend() {
// 	agentInfo := NewAgentInfo()

// }

func GetFalcoProcess() bool {
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, _ := p.Cmdline()
		if strings.Contains(cmdline, "/usr/bin/falco") {
			log.Println(cmdline)
			return true
		}
	}
	return false
}
