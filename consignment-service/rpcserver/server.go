package rpcserver

import ("github.com/micro/go-micro"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"shippy.com/consignment-service/rpcserver/api"
	"shippy.com/protoctl/consignment"
)
func Init(svc micro.Service)  {
	consignment.RegisterShippingServiceHandler(svc.Server(),new(api.ShippingService))

	if err := svc.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
		panic(err)
	}
}