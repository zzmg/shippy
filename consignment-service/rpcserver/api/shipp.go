package api

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"shippy.com/consignment-service/business"
	"shippy.com/protoctl/consignment"
)

type ShippingService struct{}

func (s *ShippingService) CreateConsignment(ctx context.Context, req *consignment.Consignment, resp *consignment.Response) error {
	err := business.CreateConsignment(req, resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *ShippingService) GetConsignment(ctx context.Context, req *consignment.GetRequest, resp *consignment.Response) error {
	err := business.GetConsignment(req, resp)
	if err != nil {
		logrus.Error(err.Error())
	}
	return nil
}
