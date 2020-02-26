package api

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"shippy.com/protoctl/consignment"
	"shippy.com/consignment-service/business"
)

type ShippingService struct {}

func (s *ShippingService)CreateConsignment(ctx context.Context,req *consignment.Consignment,resp *consignment.Response)error  {
	err := business.CreateConsignment(req,resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *ShippingService)GetConsignments(ctx context.Context,req *consignment.GetRequest,resp *consignment.Response)error {
	err := business.GetConsignments(req,resp)
	if err != nil {
		logrus.Error(err.Error())
	}
	return nil
}

