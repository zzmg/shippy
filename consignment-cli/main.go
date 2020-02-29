package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"io/ioutil"
	pb "shippy.com/protoctl/consignment"
)

func ParseFile(fileName string)(*pb.Consignment,error)  {
	data ,err := ioutil.ReadFile("consignment.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	var consignment *pb.Consignment
	err =json.Unmarshal(data,&consignment)
	if err != nil {
		fmt.Println(err.Error())
	}
	return  consignment ,nil
}

func main()  {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment",client.DefaultClient)

	consign ,_ := ParseFile("consignmengt.json")
	fmt.Println(client)
	fmt.Println(consign)
	resp, err := client.CreateConsignment(context.Background(),consign)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//resp, err = client.GetConsignments(context.Background(),&pb.GetRequest{})
	//if err != nil  {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(resp)
}
