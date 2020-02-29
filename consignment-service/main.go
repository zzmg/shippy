package main

import (
	"github.com/micro/go-micro"
	"shippy.com/consignment-service/common"
	"shippy.com/consignment-service/rpcserver"
)

func main() {
	common.LoadConfig("./conf/consignment.yaml")
	common.Initalias()
	StartService()
}

func StartService() {
	server := micro.NewService(
		micro.Name(common.GlobalConf.Micro.SvcName),
		micro.Version("latest"),
	)
	server.Init()
	rpcserver.Init(server)
}

