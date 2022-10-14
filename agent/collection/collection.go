package collection

import (
	"fmt"
	nets "net"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type AgentInfo struct {
	HostName      string
	PodName       string
	ContainerName string
	HostIp        map[string]struct{}
	ContainerIp   string
	Platform      string
	MacAddress    string
}


