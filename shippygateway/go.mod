module shippy.com/shippygateway

go 1.13

require (
	github.com/jinzhu/configor v1.1.1
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/sirupsen/logrus v1.4.2
	github.com/valyala/fasttemplate v1.1.0 // indirect
	shippy.com/protoctl v0.0.0-00010101000000-000000000000
)

replace (
	shippy.com/protoctl => ../protoctl
	shippy.com/shippygateway/common => ../../shippygateway/common
)
