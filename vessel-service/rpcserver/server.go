package rpcserver

import (
	"context"
	pb "shippy.com/protoctl/vessel"
	"shippy.com/vessel-service/business"
)

type VesselService struct {
}

func (svc *VesselService) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	business.FindAvailableVessel(spec, resp)

	return nil
}
