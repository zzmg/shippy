package common

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"shippy.com/protoctl"
	"shippy.com/vessel-service/models"
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

func ParseConf(config interface{}, path string) {
	err := configor.Load(config, path)
	if err != nil {
		msg := "Failed to load local conf file"
		logrus.Error(msg)
		panic(msg)
	}
}
func Initalise() {
	protoctl.InitLog(GlobalConf.Log)
	models.InitModel(GlobalConf.Mysql)
}

func InitLog(conf Config) {
	logrus.SetLevel(logrus.Level(conf.Log.Level))
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
