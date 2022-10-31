package dao

import "server/entity"

type Falco interface {
	Insert(entity.Falco)
	FindByIP(string) entity.Falco
	FindByAdrees(string) entity.Falco
	// FindByMacAddress(address string) entity.Device
	Update(entity.Falco)
	GetUpdateTime(string) entity.Falco
	// FindAll(updateAt time.Time) []entity.Device
	// GetUpdateTime(string) entity.Device
	// Update(device entity.Device) error
	// FindByMacAddress(address string) (*entity.Device, error)
	// FindAll(updatedAt time.Time) ([]entity.Device, error)
}
