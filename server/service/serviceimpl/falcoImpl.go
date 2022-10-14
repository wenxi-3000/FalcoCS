package serviceimpl

import (
	"server/dao"
	"server/entity"
	"server/service"
	"time"
)

type falcoService struct {
	falcoDao dao.Falco
}

func NewFalcoService(dF dao.Falco) service.FalcoService {
	return falcoService{falcoDao: dF}
}

func (d falcoService) Insert(input entity.Falco) {
	result := d.falcoDao.FindByIP(input.NodeIP)
	if result.NodeIP == "" {
		d.falcoDao.Insert(input)
	}
	d.falcoDao.Update(input)
	// log.Println(err)
}

func (d falcoService) GetUpdateTime(ip string) time.Time {
	result := d.falcoDao.GetUpdateTime(ip)
	return result.UpdatedAt
}
