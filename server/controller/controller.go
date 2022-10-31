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
	ClientService   service.ClientService
}

func NewController(
	router *gin.Engine,
	opt *libs.Options,
	deviceService service.DeviceService,
	resourceService service.ResourceService,
	falcoService service.FalcoService,
	listenner *listener.Listener,
	clientService service.ClientService,
) {
	controller := &Controller{
		DeviceService:   deviceService,
		ResourceService: resourceService,
		FalcoService:    falcoService,
		Options:         opt,
		Listenner:       listenner,
		ClientService:   clientService,
	}
	router.GET("/login", controller.getLogin)
	router.GET("/health", controller.getHealth)
	router.POST("auth", controller.postAuth)
	router.GET("/client", controller.clientHandler)
	router.POST("/device", controller.setDevice)
	router.POST("/falco", controller.setFalco)
	authGroup := router.Group("")
	authGroup.Use(libs.JWTAuthMiddleware())
	{
		authGroup.GET("/", controller.getResources)
		authGroup.GET("/devices", controller.getDevices)
		authGroup.GET("/resources/update", controller.resourcesUpdate)
		authGroup.GET("/resources", controller.getResources)
		authGroup.GET("/falco/restart", controller.restartFalco)
		authGroup.GET("/generate", controller.getGenerate)
		authGroup.POST("/generate", controller.generateClient)
		authGroup.POST("/command", controller.sendCommand)
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
