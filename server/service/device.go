package service

import (
	"server/entity"
	"time"
)

type DeviceService interface {
	//Insert(entity.Device) error
	Insert(entity.Device)
	FindAll() []entity.Device
	GetUpdateTime(string) time.Time
	FindAddressByIp(ip string) (string, error)
	FindIPByAddress(address string) (string, error)
}
