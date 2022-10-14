package sql

import (
	"log"
	"server/dao"
	"server/entity"
	"time"

	"gorm.io/gorm"
)

type deviceDao struct {
	dbClient *gorm.DB
}

// func NewDeviceDao(dbClient *gorm.DB) deviceDao {
// 	var deviceD deviceDao
// 	deviceD.dbClient = dbClient
// 	var device dao.Device
// 	return device
// }

func NewDeviceDao(dbClient *gorm.DB) dao.Device {
	return deviceDao{dbClient: dbClient}
}

func (d deviceDao) Insert(input entity.Device) {
	result := d.dbClient.Create(&input)
	log.Println(result)
}

func (d deviceDao) FindByMacAddress(address string) entity.Device {
	var device entity.Device
	log.Println("FindByMacAddress........")
	d.dbClient.Where(entity.Device{MacAddress: address}).First(&device)
	return device
}

func (d deviceDao) Update(device entity.Device) {
	// log.Println("Update..........")
	d.dbClient.First(&device)
	d.dbClient.Save(device)
	// d.dbClient.Model(&device).Where(entity.Device{MacAddress: device.MacAddress}).Update("MacAddress", device.MacAddress)
	// log.Println(device)
	// d.dbClient.Model(&device).Where(entity.Device{MacAddress: device.MacAddress}).Update(&device)
}

func (d deviceDao) FindAll(updateAt time.Time) []entity.Device {
	var devices []entity.Device
	d.dbClient.Where("updated_at > ?", updateAt.String()).Find(&devices)
	return devices
}

func (d deviceDao) GetUpdateTime(ip string) entity.Device {
	var device entity.Device
	log.Println(ip)
	d.dbClient.Where("node_ip = ?", ip).Find(&device)
	// log.Println(device.ClientIP, device.UpdatedAt)
	return device
}

// type deviceRepository struct {
// 	DbClient *gorm.DB
// }

// func NewDeviceRepository(dbClient *gorm.DB) Device {
// 	var dr deviceRepository
// 	dr.DbClient = dbClient
// 	return dr

// }

// func FindAll() {
// 	var devices []entity.Device
// 	results := dbClient.Find(devices)
// 	log.Println(results)
// }
