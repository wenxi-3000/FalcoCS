package controller

import (
	"server/libs"
	"server/listener"
	"server/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	DeviceService   service.DeviceService
	ResourceService service.ResourceService
	FalcoService    service.FalcoService
	Options         *libs.Options
	Listenner       *listener.Listener
}

func NewController(
	router *gin.Engine,
	opt *libs.Options,
	deviceService service.DeviceService,
	resourceService service.ResourceService,
	falcoService service.FalcoService,
	listenner *listener.Listener,
) {
	controller := &Controller{
		DeviceService:   deviceService,
		ResourceService: resourceService,
		FalcoService:    falcoService,
		Options:         opt,
		Listenner:       listenner,
	}

	router.POST("/device", controller.setDevice)
	router.GET("/devices", controller.getDevices)

	router.GET("/resources/update", controller.resourcesUpdate)
	router.GET("/resources", controller.getResources)
	router.POST("/falco", controller.setFalco)
	router.GET("/falco/restart", controller.restartFalco)
	router.GET("/generate", controller.getGenerate)
	router.POST("/generate", controller.generateClient)

	// router.GET("/layouts/base", func(c *gin.Context) {
	// 	log.Println("xxxxxxxxxxxx")
	// 	c.HTML(http.StatusOK, "inc/index.html", gin.H{
	// 		"title": "Posts",
	// 	})
	// })
}

// type httpController struct {
// 	DeviceService device.Service
// }

// func NewController() {
// 	handler := &httpController {
// 		DeviceService:
// 	}
// }

// func Device(c *gin.Context) {
// 	var body entity.Device
// 	c.BindJSON(&body)
// 	log.Printf("%v", &body)
// 	service.
// 	// if err := c.BindJSON(&body); err != nil {
// 	// 	log.Println(err)
// 	// 	c.Status(http.StatusBadRequest)
// 	// 	log.Printf("%v", &body)
// 	// 	return
// 	// }

// }
