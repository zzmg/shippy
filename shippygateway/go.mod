module shippy.com/shippygateway

go 1.13

require (
	github.com/jinzhu/configor v1.1.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/sirupsen/logrus v1.4.2
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/tools v0.0.0-20191029190741-b9c20aec41a5
	shippy.com/consignment-service v0.0.0-00010101000000-000000000000
	shippy.com/protoctl v0.0.0-00010101000000-000000000000
)

replace (
	shippy.com/consignment-service => ../consignment-service
	shippy.com/protoctl => ../protoctl
	shippy.com/shippygateway/common => ../shippygateway/common
)
