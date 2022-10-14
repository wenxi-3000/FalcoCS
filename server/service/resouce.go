package service

import "server/entity"

type ResourceService interface {
	//Insert(entity.Device) error
	Insert(entity.Resource)
	FindAll() []entity.Resource
}
