package dao

import (
	"server/entity"
)

type Resource interface {
	Insert(entity.Resource)
	Update(entity.Resource)
	// FindAll(updateAt time.Time) []entity.Device
	FindByIP(string) entity.Resource
	FindAll() []entity.Resource
}
