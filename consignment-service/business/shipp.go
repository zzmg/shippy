package business

import (
	"shippy.com/consignment-service/dao"
	"shippy.com/protoctl/consignment"
)

func CreateConsignment(req *consignment.Consignment, resp *consignment.Response) error {
	dao.CreateConsignment(req.ConsignmentId, req.Description, req.Weight, req.Containers, req.UserId)
	b := dao.GetConsignmentById(req.ConsignmentId)

	resp.Consignment = b
	resp.Created = true
	return nil
}

func GetConsignment(req *consignment.GetRequest, resp *consignment.Response) error {
	var all []*consignment.Consignment
	a := dao.GetConsignmentId()
	for _, v := range a {
		b := dao.GetConsignmentById(v)
		all = append(all, b)
	}

	resp.Created = true
	resp.Consignments = all
	return nil
}
