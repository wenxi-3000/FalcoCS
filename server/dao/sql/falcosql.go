package sql

import (
	"log"
	"server/dao"
	"server/entity"

	"gorm.io/gorm"
)

type falcoDao struct {
	dbClient *gorm.DB
}

// func NewDeviceDao(dbClient *gorm.DB) deviceDao {
// 	var deviceD deviceDao
// 	deviceD.dbClient = dbClient
// 	var device dao.Device
// 	return device
// }

func NewFalcoDao(dbClient *gorm.DB) dao.Falco {
	return falcoDao{dbClient: dbClient}
}

func (d falcoDao) Insert(input entity.Falco) {
	result := d.dbClient.Create(&input)
	log.Println(result)
	// result := d.dbClient.Create(&input)
	// log.Println(result)
}

func (d falcoDao) Update(input entity.Falco) {
	d.dbClient.First(&input)
	d.dbClient.Save(input)
}

func (d falcoDao) FindByIP(ip string) entity.Falco {
	var falco entity.Falco
	d.dbClient.Where(entity.Falco{NodeIP: ip}).First(&falco)
	return falco
}

func (d falcoDao) GetUpdateTime(ip string) entity.Falco {
	var falco entity.Falco
	d.dbClient.Where("node_ip = ?", ip).Find(&falco)
	// log.Println(device.ClientIP, device.UpdatedAt)
	return falco
}
