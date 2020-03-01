package main

import (
	"fmt"
	"shippy.com/shippygateway/common"
	"shippy.com/shippygateway/httpserver"
	"shippy.com/shippygateway/rpcclient"
)

func main() {
	common.LoadConfig("./conf/shippygateway.yaml")
	common.Initalias()
	fmt.Println("start rpc client......")
	rpcclient.StartClient()
	fmt.Println("end  rpc client......")
	fmt.Println("start  http server ......")
	httpserver.StartServer()
	fmt.Println("end  http server ......")
}

