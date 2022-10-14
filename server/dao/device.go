package dao

import (
	"server/entity"
	"time"
)

type Device interface {
	Insert(entity.Device)
	FindByMacAddress(address string) entity.Device
	Update(device entity.Device)
	FindAll(updateAt time.Time) []entity.Device
	GetUpdateTime(string) entity.Device
	// Update(device entity.Device) error
	// FindByMacAddress(address string) (*entity.Device, error)
	// FindAll(updatedAt time.Time) ([]entity.Device, error)
}
