package  main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"shippy.com/protoctl/consignment"
	vesselpb "shippy.com/protoctl/vessel"
)

type Repository struct {
	consignments []*consignment.Consignment
}

type IRepository interface {
	Create(consignment *consignment.Consignment)(*consignment.Consignment,error)
}

func (repo *Repository)Create(con *consignment.Consignment)(*consignment.Consignment,error)  {
	repo.consignments = append(repo.consignments,con)
	return con,nil
}
func (repo *Repository)GetAll()([] *consignment.Consignment){
	return repo.consignments
}
type ShippingService struct {
	repo Repository
	vesselClient vesselpb.VesselServiceClient
}
//gRpc
//func (s *ShippingService)CreateConsignment(ctx context.Context,con *consignment.Consignment) (*consignment.Response,error) {
//	curennt_consign ,err := s.repo.Create(con)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	resp := &consignment.Response{
//		Created:              true,
//		Consignment:          curennt_consign,
//	}
//	return resp,nil
//}
//go-micro

func (s *ShippingService)CreateConsignment(ctx context.Context,req *consignment.Consignment,resp *consignment.Response)error  {
	fmt.Println("going into service.......")
	current_consign ,err := s.repo.Create(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Created = true
	resp.Consignment = current_consign
	fmt.Println(resp)

	var vReq vesselpb.Specification
	vReq.Capacity = int32(len(req.Containers))
	vReq.MaxWeight= req.Weight
	vResp ,err := s.vesselClient.FindAvailable(context.Background(),&vReq)
	if err != nil {
		return err
	}
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	con,err := s.repo.Create(req)
	if err != nil {
		return  err
	}
	resp.Created = true
	resp.Consignment = con
	return nil
}

//func (s *ShippingService)GetConsignments(ctx context.Context ,con *consignment.GetRequest) (*consignment.Response,error)  {
//	AllConsignments := s.repo.GetAll()
//	resp := &consignment.Response{
//		Consignments:         AllConsignments,
//	}
//	return resp,nil
//}

func (s *ShippingService)GetConsignments(ctx context.Context,req *consignment.GetRequest,resp *consignment.Response)error {
	AllConsignment := s.repo.GetAll()
	resp = &consignment.Response{
		Created:              true,
		Consignments:         AllConsignment,
	}
	fmt.Println(resp)
	return  nil
}
func main()  {
	//listener,err := net.Listen("tcp",":8888")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//log.Printf("listen on: 8888\n")
	//server := grpc.NewServer()
	////consignment.RegisterShippingServiceServer(server,new(ShippingService))
	//consignment.RegisterShippingServiceHandler(server,)
	//if err :=server.Serve(listener);err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init()
	var repo Repository
	vClient := vesselpb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	consignment.RegisterShippingServiceHandler(server.Server(), &ShippingService{repo,vClient})
	if err := server.Run() ;err != nil {
		fmt.Println("failed to serve: %v", err)
	}

}