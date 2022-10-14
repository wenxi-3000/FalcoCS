package serviceimpl

import (
	"server/dao"
	"server/entity"
	"server/service"
)

type resourceService struct {
	resourceDao dao.Resource
}

func NewResourceService(daor dao.Resource) service.ResourceService {
	return resourceService{resourceDao: daor}
}

func (d resourceService) Insert(input entity.Resource) {
	resource := d.resourceDao.FindByIP(input.NodeIP)
	if resource.NodeIP == "" {
		d.resourceDao.Insert(input)
	}
	d.resourceDao.Update(input)
}

func (d resourceService) FindAll() []entity.Resource {
	results := d.resourceDao.FindAll()
	return results
}
