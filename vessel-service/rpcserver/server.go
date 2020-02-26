package rpcserver

import (
	"context"
	"errors"
	"fmt"
	pb "shippy.com/protoctl/vessel"
)
type Repository interface {
	FindAvailable(*pb.Specification)(*pb.Vessel,error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

type VesselService struct {
	repo Repository
}

func (ves * VesselRepository)FindAvailable(spec *pb.Specification)(*pb.Vessel,error)  {
	for _,v := range ves.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v,nil
		}
	}
	return nil, errors.New("No vessel can't be use")
}

func(svc *VesselService)FindAvailable(ctx context.Context,spec *pb.Specification,resp *pb.Response)error  {
	v ,err := svc.repo.FindAvailable(spec)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Vessel = v
	return nil
}
