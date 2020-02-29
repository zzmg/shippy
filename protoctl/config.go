package protoctl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type ConfigService struct {
	SvcName   string   `yaml:"svc_name"`
	SvcAddr   string   `yaml:"svc_addr"`
	EtcdAddrs []string `yaml:"etcd_addrs"`
}
type ConfigMysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxConn  int    `yaml:"max_conn"`
	Charset  string `yaml:"charset"`
}

type ConfigLog struct {
	Level      int    `yaml:"level"`
	Path       string `yaml:"path"`
	OutputDest string `yaml:"output_dest"`
}

func CreateDB(config ConfigMysql) *gorm.DB {
	var (
		username = config.Username
		password = config.Password
		host     = config.Host
		port     = config.Port
		dbName   = config.DBName
		charset  = config.Charset
		maxIdle  = config.MaxIdle
		maxOpen  = config.MaxConn
	)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		dbName,
		charset,
	)
	fmt.Println(connStr)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect MYSQL %s:%d/%s: %s", host, port, dbName, err.Error()))
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.AutoMigrate()
	return db
}

func InitLog(conf ConfigLog) {
	logrus.SetLevel(logrus.Level(conf.Level))
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
