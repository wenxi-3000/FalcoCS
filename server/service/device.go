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
}
