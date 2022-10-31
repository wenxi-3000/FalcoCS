package connect

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type Falco struct {
	ClientID string `json:"client_id,omitempty"`
	Falco    bool   `json:"falco"`
}

func (c *connector) NewFalco() ([]byte, error) {
	var falco Falco

	falco.ClientID = c.clientId
	var err error
	falco.Falco, err = GetFalcoProcess()
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(falco)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func InfoSend() {
// 	agentInfo := NewAgentInfo()

// }

func GetFalcoProcess() (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}
	for _, p := range processes {
		// log.Println(p.Cmdline())
		cmdline, _ := p.Cmdline()
		if strings.Contains(cmdline, "/usr/bin/falco") {
			log.Println(cmdline)
			return true, nil
		}
	}
	return false, nil
}

func (c *connector) SendFalco() error {
	falco, err := c.NewFalco()
	if err != nil {
		return err
	}
	url := c.rootUri + "/falco"
	resp, err := c.NewRequest(http.MethodPost, url, falco)
	if err != nil {
		return err
	}

	if resp.code != http.StatusOK {
		return fmt.Errorf("error with status code %d", resp.code)
	}
	return nil

	// var hasError bool
	// falco, err := GetFalcoProcess()
	// if err != nil {
	// 	hasError = true
	// 	return err
	// }

	// body, erra := json.Marshal(&CmdResponse{
	// 	ClientID: c.clientId,
	// 	Falco:    falco,
	// 	HasError: hasError,
	// })

	// if erra != nil {
	// 	hasError = true
	// 	return erra
	// }
	// errs := c.websockconn.WriteMessage(websocket.BinaryMessage, body)
	// if errs != nil {
	// 	return errs

	// }
}
