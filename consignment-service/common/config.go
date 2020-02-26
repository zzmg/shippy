package common

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"shippy.com/consignment-service/models"
)
var (
	GlobalConf *Config
)
type Config struct {
	Micro ConfigService `yaml:"micro"`
	Mysql ConfigMysql	`yaml:"mysql"`
	Log ConfigLog `yaml:"log"`
}
type ConfigService struct {
	SvcName          string   `yaml:"svc_name"`
	SvcAddr          string   `yaml:"svc_addr"`
	EtcdAddrs  []string `yaml:"etcd_addrs"`
}
type ConfigMysql struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	DBName         string `yaml:"db_name"`
	MaxIdle        int    `yaml:"max_idle"`
	MaxConn        int    `yaml:"max_conn"`
	Charset		string `yaml:"charset"`
}

type ConfigLog struct {
	Level      int      `yaml:"level"`
	Path       string   `yaml:"path"`
	OutputDest string   `yaml:"output_dest"`
}

func LoadConfig(filePath string){
	GlobalConf := &Config{}
	ParseConf(GlobalConf,filePath)
}

func ParseConf(config interface{},path string)  {
	err := configor.Load(config , path)
	if err != nil {
		msg := "Failed to load local conf file"
		logrus.Error(msg)
		panic(msg)
	}
}
func Initalise() {
	InitLog(GlobalConf.Log)
	models.InitModel(GlobalConf.Mysql)
}

func InitLog(conf ConfigLog) {
	logrus.SetLevel(logrus.Level(conf.Level))
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
