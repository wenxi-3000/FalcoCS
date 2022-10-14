package heartbeat

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type AgentInfo struct {
	hostName      string
	podName       string
	containerName string
	hostIp        string
	containerIp   string
	platformInfo  string
	falcoInfo
}

func GetAgentStat() {
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Println(v)
}

func GetHostInfo() *host.InfoStat {
	hostInfo, _ := host.Info()
	return hostInfo
}

func GetHostName() {
	hostInfo, _ := host.Info()
	hostName := hostInfo.Hostname
	fmt.Println(hostName)
}

func GetPlatform() {
	hostInfo, _ := host.Info()
	hostPlatform := hostInfo.Platform
	fmt.Println(hostPlatform)
}

func GetProcess() {
	processes, _ := process.Processes()
	for _, p := range processes {
		fmt.Println(p.Cmdline())
	}
}

func GetNetInfo() {
	// var netConnectionKindMap = map[string][]netConnectionKindType{
	// 	"all":   {kindTCP4, kindTCP6, kindUDP4, kindUDP6, kindUNIX},
	// 	"tcp":   {kindTCP4, kindTCP6},
	// 	"tcp4":  {kindTCP4},
	// 	"tcp6":  {kindTCP6},
	// 	"udp":   {kindUDP4, kindUDP6},
	// 	"udp4":  {kindUDP4},
	// 	"udp6":  {kindUDP6},
	// 	"unix":  {kindUNIX},
	// 	"inet":  {kindTCP4, kindTCP6, kindUDP4, kindUDP6},
	// 	"inet4": {kindTCP4, kindUDP4},
	// 	"inet6": {kindTCP6, kindUDP6},
	// }

	hostIpx, _ := net.Connections("inet4")
	for _, ip := range hostIpx {
		fmt.Println(ip.Laddr.IP)
	}

}
