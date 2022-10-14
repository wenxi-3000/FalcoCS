package sql

import (
	"server/dao"
	"server/entity"

	"gorm.io/gorm"
)

type resourceDao struct {
	dbClient *gorm.DB
}

// func NewDeviceDao(dbClient *gorm.DB) deviceDao {
// 	var deviceD deviceDao
// 	deviceD.dbClient = dbClient
// 	var device dao.Device
// 	return device
// }

func NewResourceDao(dbClient *gorm.DB) dao.Resource {
	return resourceDao{dbClient: dbClient}
}

func (d resourceDao) Insert(input entity.Resource) {
	d.dbClient.Create(&input)
}

func (d resourceDao) Update(input entity.Resource) {
	d.dbClient.First(&input)
	d.dbClient.Save(&input)
}

func (d resourceDao) FindByIP(ip string) entity.Resource {
	var resource entity.Resource
	d.dbClient.Where(entity.Resource{NodeIP: ip}).First(&resource)
	return resource
}

func (d resourceDao) FindAll() []entity.Resource {
	var resources []entity.Resource
	d.dbClient.Find(&resources)
	return resources
}
