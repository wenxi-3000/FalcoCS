package service

import (
	"server/entity"
	"time"
)

type FalcoService interface {
	//Insert(entity.Device) error
	Insert(entity.Falco)
	GetUpdateTime(string) time.Time
}
