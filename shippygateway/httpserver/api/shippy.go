package api

import (
	"context"
	"github.com/labstack/echo"
	pb "shippy.com/protoctl/consignment"
	"shippy.com/shippygateway/httpserver/helper"
	"shippy.com/shippygateway/rpcclient"
)

type consignmentArgs struct {
	ConsignmentId string          `json:"id" required:"true"`
	Description   string          `json:"description" required:"true"`
	Weight        int32           `json:"weight" required:"true"`
	Containers    []*pb.Container `json:"container_id" required:"true"`
	UserId        string          `json:"user_id" required:"true"`
}

/*
 @ConsignmentId  创建货物的id
 @Description    货物的描述
 @Weight		 货物的重量
 @ContainerId    货物放在集装箱的id
 @UserId         货物的owner
 @Router /v1/shippy/consignment/create [post]
*/
func CreateConsignment(ctx echo.Context) error {
	args := &consignmentArgs{}
	if err := ctx.Bind(args); err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	var pbcon pb.Consignment
	pbcon.ConsignmentId = args.ConsignmentId
	pbcon.UserId = args.UserId
	pbcon.Weight = args.Weight
	pbcon.Description = args.Description
	pbcon.Containers = args.Containers
	rpcResp, err := rpcclient.ShippyServiceClient.CreateConsignment(context.Background(), &pbcon)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, rpcResp)
}

func GetConsignment(ctx echo.Context) error {
	rpcResp, err := rpcclient.ShippyServiceClient.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, rpcResp)
}
