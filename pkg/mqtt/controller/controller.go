package controller

import "Project/pkg/mqtt/storage"

var Controller SensorController

func Init() {
	Controller = SensorController{storage.Dao}
}
