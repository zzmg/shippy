package models

import (
	"github.com/jinzhu/gorm"
	"shippy.com/consignment-service/common"
	"shippy.com/consignment-service/models"
)

var db *gorm.DB

func DB()  *gorm.DB {
	return db
}

func InitModel(config common.ConfigMysql)  {
	db = models.CreateDB(config)
	db.AutoMigrate(
		&Vessel{},
			)
}
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
type Vessel struct {
	Id string
	Capacity int32
	MaxWeight int32
	Name string
	Available bool
	OwnerId string
	Start string
	End string
}

