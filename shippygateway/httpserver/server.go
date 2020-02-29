package httpserver

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	serviceapi "shippy.com/shippygateway/httpserver/api"
)

var api = echo.New()

func init() {
	api.Use(middleware.Logger())
	api.Use(middleware.CORS())
	api.Use(middleware.Recover())
	MountAPIModule(api)
}

func MountAPIModule(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	MountAPI(apiv1)
}

// MountAPI registers API
func MountAPI(group *echo.Group) {
	MountShippyAPI(group)
}

func MountShippyAPI(group *echo.Group) {
	task := group.Group("/shippy")
	task.POST("/consignment/create", serviceapi.CreateConsignment)
	task.GET("/consignment/getall", serviceapi.GetConsignment)
}

func StartServer() {
	api.Start("18888")
}
