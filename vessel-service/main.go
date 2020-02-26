package main

import (
	"github.com/micro/go-micro"
	pb "shippy.com/protoctl/vessel"
)

func main()  {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
		)
	server.Init()
	pb.RegisterVesselServiceHandler(server.Server(),&VesselService{repo})

	if err:= server.Run();err != nil {
		fmt.Println(err.Error())
	}
}