package main

import (
	"fmt"
	"server/controller"
	"server/dao"
	"server/dao/sql"
	"server/libs"
	"server/listener"
	"server/service/serviceimpl"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("==========FalcoCS Start==========")
	opt := libs.NewOptions()
	//初始化数据库
	opt.DB = dao.InitDB(*opt)

	//tcp连接池
	listener := listener.Newlistener(libs.ServerIP, libs.ServerPort)

	//初始化设备业务层，存储层
	deviceDao := sql.NewDeviceDao(opt.DB)
	deviceService := serviceimpl.NewDeviceService(deviceDao)

	//初始化资源业务层，存储层
	resourceDao := sql.NewResourceDao(opt.DB)
	resourceService := serviceimpl.NewResourceService(resourceDao)

	//初始化Falco业务层，存储层
	falcoDao := sql.NewFalcoDao(opt.DB)
	falcoService := serviceimpl.NewFalcoService(falcoDao)

	//初始化client
	clientService := serviceimpl.NewClientService(opt)
	//初始化device
	// deviceDao := sql.NewDeviceDao(opt.DB)
	// deviceService := serviceimpl.NewDeviceService(deviceDao)
	// log.Println(deviceService)
	// log.Println(deviceService.FindAll())

	// log.Println(deviceService.Insert())

	//初始化gin资源
	router := gin.Default()
	router.Static("/static", "web/static")
	// LoadTemplates("web/templates")
	// router.LoadHTMLGlob("web/templates/**/*")
	router.HTMLRender = libs.LoadTemplates("web/templates")
	// router.POST("/device", controller.SetDeviceHandler)
	controller.NewController(router, opt, deviceService, resourceService, falcoService, listener, clientService)
	// router.POST("/collection", func(c *gin.Context) {
	// 	json := PostJsonData{}
	// 	c.BindJSON(&json)
	// 	log.Printf("%v", &json)
	// })

	// //

	// router.POST("/devices", controller.Device)
	// router.GET("/posts/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	// 		"title": "Posts",
	// 	})
	// })

	router.Run(":8081")

}

// func Insert(input entities.Device) error {
// 	_, err := d.Repository.FindByMacAddress(input.MacAddress)
// 	if errors.Is(err, repositories.ErrNotFound) {
// 		return d.Repository.Insert(input)
// 	}
// 	return d.Repository.Update(input)
// }
