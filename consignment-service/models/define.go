package models

import (
	_ "github.com/go-sql-driver/mysql"
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
		&Container{},
		&Consignment{})
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

type Container struct {
	ContainerId string `json:"container_id"`
	CustomerId  string `json:"customer_id"`
	UserId      string `json:"user_id"`
}

type Consignment struct {
	ConsignmentId string `json:"id"`
	Description   string `json:"description"`
	Weight        int32  `json:"weight"`
	ContainerId   string `json:"container_id"`
	UserId        string `json:"user_id"`
}
