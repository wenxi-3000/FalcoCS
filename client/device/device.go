package device

import (
	"encoding/json"
	"fmt"
	"log"
	nets "net"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Device struct {
	Hostname   string   `json:"hostname"`
	MacAddress string   `json:"mac_address"`
	IPs        []string `json:"ips"`
	// Falco      bool     `json:"falco"`
}

func NewDevice() []byte {
	var device Device
	hostInfo := GetHostInfo()
	device.Hostname = hostInfo.Hostname
	device.IPs = GetLaddrIP()
	device.MacAddress, _ = GetMacAddress()
	// device.Falco = GetProcess()

	result, err := json.Marshal(device)
	if err != nil {
		log.Println(err)
	}
	return result
}

// func InfoSend() {
// 	agentInfo := NewAgentInfo()

// }

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

// func GetProcess() bool {
// 	processes, _ := process.Processes()
// 	for _, p := range processes {
// 		cmdline, _ := p.Cmdline()
// 		if strings.Contains(cmdline, "/usr/bin/falco") {
// 			log.Println(cmdline)
// 			return true
// 		}
// 	}
// 	return false
// }

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

//返回非127.0.0.1和0.0.0.0的ipv4地址
func GetLaddrIP() []string {
	var results []string
	inet4ips := map[string]struct{}{}
	ips, _ := net.Connections("tcp4")
	for _, ip := range ips {
		if ip.Laddr.IP != "127.0.0.1" && ip.Laddr.IP != "0.0.0.0" {
			// fmt.Println(ip.Laddr.IP)
			inet4ips[ip.Laddr.IP] = struct{}{}
		}
	}
	//fmt.Println(inet4ips)
	for result := range inet4ips {
		results = append(results, result)
	}
	return results
	// for item := range inet4ips {
	// 	fmt.Println(item)
	// }
	// return inet4ips
}

func GetMacAddress() (string, error) {
	interfaces, err := nets.Interfaces()
	if err != nil {
		return "", err
	}
	var address []string
	for _, i := range interfaces {
		a := i.HardwareAddr.String()
		if a != "" {
			address = append(address, a)
		}
	}
	if len(address) == 0 {
		return "", nil
	}
	return address[0], nil
}
