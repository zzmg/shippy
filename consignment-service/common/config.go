package common

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"shippy.com/consignment-service/models"
	"shippy.com/protoctl"
)

var (
	GlobalConf *Config
)

type Config struct {
	Micro protoctl.ConfigService `yaml:"micro"`
	Mysql protoctl.ConfigMysql   `yaml:"mysql"`
	Log   protoctl.ConfigLog     `yaml:"log"`
}

func LoadConfig(filePath string) {
	GlobalConf = &Config{}
	ParseConf(GlobalConf, filePath)
}

func Initalias() {
	protoctl.InitLog(GlobalConf.Log)
	models.InitModel(GlobalConf.Mysql)
}

func ParseConf(dest interface{}, path string) {
	err := configor.Load(dest, path)
	if err != nil {
		msg := "Failed to load local conf file"
		logrus.Error(err.Error())
		panic(msg)
	}
}
