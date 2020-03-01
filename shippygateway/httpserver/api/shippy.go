package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	pb "shippy.com/protoctl/consignment"
	"shippy.com/shippygateway/httpserver/helper"
	"shippy.com/shippygateway/rpcclient"
)

/*
 @ConsignmentId  创建货物的id
 @Description    货物的描述
 @Weight		 货物的重量
 @ContainerId    货物放在集装箱的id
 @UserId         货物的owner
 @Router /v1/shippy/consignment/create [post]
*/
func CreateConsignment(ctx echo.Context) error {
	args := new(pb.Consignment)

	if err := ctx.Bind(args); err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	fmt.Println("ddddddddd")
	rpcResp, err := rpcclient.ShippyServiceClient.CreateConsignment(context.Background(), args)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, rpcResp)
}

/*
 @Router /v1/consignment/getall [get]
*/

func GetConsignment(ctx echo.Context) error {
	rpcResp, err := rpcclient.ShippyServiceClient.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, rpcResp)
}
