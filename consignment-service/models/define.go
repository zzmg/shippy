package models

import (
	"github.com/jinzhu/gorm"
	"shippy.com/consignment-service/common"
	"fmt"
	"shippy.com/protoctl/consignment"
)

var db *gorm.DB

func DB()  *gorm.DB {
	return db
}

func InitModel(config common.ConfigMysql)  {
	db = CreateDB(config)
	db.AutoMigrate(
		&Container{},
		&Consignment{},)
}

func CreateDB(config common.ConfigMysql)  *gorm.DB {
	var (
		username = config.Username
		password = config.Password
		host     = config.Host
		port     = config.Port
		dbName   = config.DBName
		maxIdle  = config.MaxIdle
		maxOpen  = config.MaxConn
	)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		dbName,
		config.Charset,
	)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect MYSQL %s:%d/%s: %s", host, port, dbName, err.Error()))
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.AutoMigrate()

	return db
}
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

type Container struct {
	Id string
	CustomerId  string
	UserId	string
	Start string
	End string
}

type Consignment struct {
	Id string
	Description string
	Weight int32
	Containers []*Container
	VesselId string
}