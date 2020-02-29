package dao

import (
	"github.com/sirupsen/logrus"
	consignmentmodels "shippy.com/consignment-service/models"
	"shippy.com/protoctl/vessel"
	"shippy.com/vessel-service/models"
)

const ContainersCapacity = 100

func CreateVessel(id string, capacity int32, maxWeight int32, name string, available bool, ownerId string, start string, end string) *models.Vessel {
	var ves models.Vessel
	{
		ves.Id = id
		ves.Capacity = capacity
		ves.MaxWeight = maxWeight
		ves.Name = name
		ves.Available = available
		ves.OwnerId = ownerId
		ves.Start = start
		ves.End = end
	}
	db := models.DB().Find(&ves)
	if err := db.Error; err != nil {
		logrus.Error("mysql", err, "failed to create vessel")
	}
	return &ves
}

func FindAvailableVessel(req *vessel.Specification) []*models.Vessel {
	maxWeight := req.MaxWeight
	capacity := req.Capacity

	db := consignmentmodels.DB()

	return nil
}
