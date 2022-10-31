package connect

import (
	"fmt"
	"log"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type FalcoInfo struct {
	Pid     int32
	running bool
	cmdline string
}

func RestartFalco() (bool, error) {
	log.Println("Restart Falco")

	//获取进程信息
	falcoInfo, err := GetFalcoInfo()
	if err != nil {
		log.Println(err)
		if fmt.Sprint(err) == "Falco Processes Not Found" {
			success, err := StartFalco()
			if err != nil {
				return false, err
			}
			log.Println(success, err)
			return success, nil
		}
	}

	if falcoInfo.cmdline != "" {
		//杀死Falco进程
		KillFalco(falcoInfo.Pid)
		//重启Falco进程
		success, err := StartFalco()
		if err != nil {
			return false, err
		}
		return success, nil
	}

	return false, nil

}

//重启Falco进程
func StartFalco() (bool, error) {
	cmd := "nohup /usr/bin/falco &"
	output, err := RunCommandWithErr(cmd)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println(output)
	return true, nil

}

//杀死Falco进程
func KillFalco(pid int32) {
	cmd := "kill -9 " + fmt.Sprint(pid)
	log.Println(cmd)
	output, err := RunCommandWithErr(cmd)
	if err != nil {
		log.Println(err)
	}
	log.Println(output)
}

func GetFalcoInfo() (*FalcoInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, p := range processes {
		// log.Println(p.Cmdline())
		cmdline, _ := p.Cmdline()

		if strings.Contains(cmdline, "/usr/bin/falco") {
			log.Println(p.Pid)
			log.Println(cmdline)
			return &FalcoInfo{
					Pid:     p.Pid,
					running: true,
					cmdline: cmdline,
				},
				nil
		}
	}

	return nil, fmt.Errorf("Falco Processes Not Found")
}
