package dao

import (
	"github.com/sirupsen/logrus"
	"shippy.com/consignment-service/models"
	"shippy.com/protoctl/consignment"
)

func CreateConsignment(consgnmentId string, description string, weight int32, containers []*consignment.Container, userId string) []*models.Consignment {
	var con models.Consignment
	{
		con.ConsignmentId = consgnmentId
		con.Description = description
		con.Weight = weight
		con.UserId = userId
	}
	var allcon []*models.Consignment
	var container models.Container
	for _, v := range containers {
		con.ContainerId = v.ContainerId
		db := models.DB().Create(&con)
		container.UserId = v.UserId
		container.CustomerId = v.CustomerId
		container.ContainerId = v.ContainerId
		db = models.DB().Create(&container)
		if err := db.Error; err != nil {
			logrus.Error("mysql", err, "failed to create consignment")
		}
	}
	models.DB().Where("user_id = ?", userId).Find(&allcon)
	return allcon
}

func GetConsignment() []*models.Consignment {
	var allConsignment []*models.Consignment
	models.DB().Find(allConsignment)
	return allConsignment
}

type ResultId struct {
	ConsignmentId string
}

func GetConsignmentId() []*ResultId {
	var allId []*ResultId
	models.DB().Table("consignments").Select("DISTINCT(consignment_id)").Find(&allId)
	return allId
}
func GetConsignmentById(consignment_id string) *consignment.Consignment {
	var a []models.Consignment
	var protocon consignment.Consignment
	models.DB().Where("consignment_id = ?", consignment_id).Find(&a)
	for _, v := range a {
		var container_model models.Container
		var container_proto consignment.Container
		container_id := v.ContainerId
		models.DB().Table("containers").Where("container_id = ?", container_id).Find(&container_model)
		container_proto.ContainerId = container_model.ContainerId
		container_proto.UserId = container_model.UserId
		container_proto.CustomerId = container_model.CustomerId
		protocon.ConsignmentId = v.ConsignmentId
		protocon.Description = v.Description
		protocon.UserId = v.UserId
		protocon.Weight = v.Weight
		protocon.Containers = append(protocon.Containers, &container_proto)
	}
	return &protocon
}
