package libs

import (
	"log"
	"server/utils"

	"gorm.io/gorm"
)

type Options struct {
	DBType        string
	DBPath        string
	DB            *gorm.DB
	Resources     []Resources
	NodeIPs       []string
	ClusterNames  []string
	ServerPort    string
	ServerAddress string
	ClientName    string
}

type ReceiveClient struct {
	IPs        []string `json:"ips"`
	Hostname   string   `json:"hostname"`
	MacAddress string   `json:"mac_address"`
}

func NewOptions() *Options {
	//日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var opt Options
	opt.DBType = DatabaseType
	opt.DBPath = DatabaseDirectory + "/" + DatabaseName
	opt.ServerPort = ServerPort
	opt.ServerAddress = ServerIP
	opt.ClientName = ClientName
	//初始化目录
	if err := MakeDir(TmpDirectory, DatabaseDirectory); err != nil {
		log.Println(err)
	}

	opt.Resources = ParseResources(ResourcesPath)
	for _, resource := range opt.Resources {
		opt.ClusterNames = append(opt.ClusterNames, resource.Name)
		opt.NodeIPs = append(opt.NodeIPs, resource.IP...)
	}
	return &opt
}

func MakeDir(paths ...string) error {
	for _, path := range paths {
		resultPath, err := utils.NormalizePath(path)
		if err != nil {
			return err
		}
		if !utils.FolderExists(resultPath) {
			utils.MakeDir(resultPath)
		}
	}
	return nil
}
