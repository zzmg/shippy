module shippy.com/consignment-service

go 1.13

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/configor v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/tools v0.0.0-20191029190741-b9c20aec41a5
	google.golang.org/grpc v1.25.1
	shippy.com/protoctl v0.0.0-00010101000000-000000000000
)

replace (
	gitlab.wallstcn.com/wscnbackend/ivankastd => ../../gitlab.wallstcn.com/wscnbackend/ivankastd
	shippy.com/protoctl => ../protoctl
)
