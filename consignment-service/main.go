package  main

import (
	"github.com/micro/go-micro"
	"shippy.com/consignment-service/common"
	"shippy.com/consignment-service/rpcserver"
)

func main()  {
	common.LoadConfig("./conf/consignment.yaml")
	common.Initalise()
	StartService(common.ConfigService{})
}

func StartService(config common.ConfigService) {
	server := micro.NewService(
		micro.Name(config.SvcName),
		micro.Version("latest"),
	)
	server.Init()
	rpcserver.Init(server)
}