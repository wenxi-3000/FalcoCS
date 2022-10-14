package dao

import (
	"log"
	"server/entity"
	"server/libs"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(opt libs.Options) *gorm.DB {
	var err error
	//使用sqlite
	if opt.DBType == "sqlite" && opt.DBPath != "" {
		DB, err = gorm.Open(sqlite.Open(opt.DBPath))
		if err != nil {
			log.Println(err)
		}
		DB.AutoMigrate(&entity.Device{})
		DB.AutoMigrate(&entity.Resource{})
		DB.AutoMigrate(&entity.Falco{})
	}
	return DB
}
