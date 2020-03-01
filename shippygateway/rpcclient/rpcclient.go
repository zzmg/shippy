package rpcclient

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/sirupsen/logrus"
	pb "shippy.com/protoctl/consignment"
	"shippy.com/shippygateway/common"
)

var ShippyServiceClient pb.ShippingServiceClient

func StartClient() {
	server := micro.NewService(
		micro.Name(common.GlobalConf.Micro.SvcName),
	)
	server.Init()
	ShippyServiceClient = pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				logrus.Error(e)
			}
		}()
		server.Run()
	}()

}
