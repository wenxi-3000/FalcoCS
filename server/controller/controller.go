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
	router.GET("/login", controller.getLogin)
	router.POST("auth", controller.postAuth)
	authGroup := router.Group("")
	authGroup.Use(libs.JWTAuthMiddleware())
	{
		authGroup.GET("/", controller.getResources)
		authGroup.POST("/device", controller.setDevice)
		authGroup.GET("/devices", controller.getDevices)
		authGroup.GET("/resources/update", controller.resourcesUpdate)
		authGroup.GET("/resources", controller.getResources)
		authGroup.POST("/falco", controller.setFalco)
		authGroup.GET("/falco/restart", controller.restartFalco)
		authGroup.GET("/generate", controller.getGenerate)
		authGroup.POST("/generate", controller.generateClient)
	}

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
