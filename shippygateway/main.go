package main

import (
	"shippy.com/shippygateway/common"
	"shippy.com/shippygateway/httpserver"
	"shippy.com/shippygateway/rpcclient"
)

func main() {
	common.LoadConfig("./conf/shippygateway.yaml")
	common.Initalias()
	httpserver.StartServer()
	rpcclient.StartClient()
}

