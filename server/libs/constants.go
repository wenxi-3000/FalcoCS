package libs

import (
	"time"
)

const (
	TimeoutDuration     = time.Second * 30
	TmpDirectory        = "tmp/"
	DatabaseDirectory   = "database"
	DatabaseName        = "FalcoCS.db"
	DatabaseType        = "sqlite"
	AppName             = "FalcoCS"
	ResourcesPath       = "conf/resources.yaml"
	ServerIP            = "172.16.42.150"
	ServerPort          = "8081"
	ClientName          = "falcoc"
	LoginUsername       = "admin"
	LoginPassword       = "admin"
	JwtSecret           = "ShadowFl0w"
	TokenExpireDuration = time.Hour * 24
)
