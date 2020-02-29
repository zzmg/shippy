package models

import (
	"github.com/jinzhu/gorm"
	"shippy.com/protoctl"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func InitModel(config protoctl.ConfigMysql) {
	db = protoctl.CreateDB(config)
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
	VesselId  string
	Capacity  int32
	MaxWeight int32
	Name      string
	Available bool
	OwnerId   string
}
