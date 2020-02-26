package dao

import (
	"github.com/sirupsen/logrus"
	"shippy.com/consignment-service/models"
)

func CreateConsignment(id string,description string, weight int32,containers []*models.Container,vesselId string)(*models.Consignment)  {
	var con models.Consignment
	{
		con.Id = id
		con.Description = description
		con.Weight = weight
		con.Containers = containers
		con.VesselId = vesselId
	}
	db := models.DB().Create(&con)
	if err := db.Error; err != nil {
		logrus.Error("mysql", err, "failed to create consignment")
	}
	return &con
}
func GetConsignments()([]*models.Consignment){
	var allConsignment []*models.Consignment
	models.DB().Find(allConsignment)
	return allConsignment
}
