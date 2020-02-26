package business

import (
	"shippy.com/protoctl/consignment"
	"shippy.com/consignment-service/dao"
	"shippy.com/consignment-service/models"
)

func CreateConsignment(req *consignment.Consignment,resp *consignment.Response) error {
	conn := PbConsignmentToModelDb(req)
	con  := dao.CreateConsignment(conn.Id,conn.Description,conn.Weight,conn.Containers,conn.VesselId)
	r := ModelDbConsignmentToPb(con)
	resp.Created = true
	resp.Consignment = r
	return nil
}

func GetConsignments(req *consignment.GetRequest,resp *consignment.Response) error  {
	allConsignment := dao.GetConsignments()
	var allPb []*consignment.Consignment
	for _,v := range  allConsignment {
		pb := ModelDbConsignmentToPb(v)
		allPb = append(allPb,pb)
	}
	resp.Created = true
	resp.Consignments = allPb
	return  nil
}

func ModelDbConsignmentToPb(req *models.Consignment) (*consignment.Consignment) {
	var a consignment.Consignment
	a.Description = req.Description
	a.VesselId = req.VesselId
	a.Id = req.Id
	a.Weight = req.Weight
	var b []*consignment.Container
	for _,v := range req.Containers {
		var c consignment.Container
		c.Id = v.Id
		c.CustomerId = v.CustomerId
		c.UserId = v.UserId
		c.Start = v.Start
		c.End = c.End
		b = append(b,&c)
	}
	a.Containers = b
	return &a
}
func PbConsignmentToModelDb(req *consignment.Consignment) (*models.Consignment) {
	var a models.Consignment
	a.Description = req.Description
	a.VesselId = req.VesselId
	a.Id = req.Id
	a.Weight = req.Weight
	var b  []*models.Container
	for _,v := range req.Containers {
		var c models.Container
		c.Id = v.Id
		c.CustomerId = v.CustomerId
		c.UserId = v.UserId
		c.Start = v.Start
		c.End = c.End
		b = append(b,&c)
	}
	a.Containers = b
	return &a
}

