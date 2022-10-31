package controller

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"server/entity"
	"server/libs"
	"server/service"
	"server/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ct *Controller) getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (ct *Controller) postAuth(c *gin.Context) {
	var user libs.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	if user.Username == ct.Options.UserInfo.Username && user.Password == ct.Options.UserInfo.Password {
		// 生成Token
		tokenString, err := libs.GenToken(user.Username)
		log.Println(c.Request.Host)
		c.SetCookie("Authorization", tokenString, 3600, "/", "", false, true)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})

}

func (ct *Controller) setDevice(c *gin.Context) {
	var body libs.ReceiveClient
	var entityDevice entity.Device
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
	}
	log.Println(body)

	//客户端获取到的Laddr ip 只有一个
	hostIP, err := utils.GetHostIp(body.IPs, ct.Options.NodeIPs)
	if err != nil {
		log.Println(err)
	}
	entityDevice.ClientIPS = utils.SliceToString(body.IPs)
	entityDevice.NodeIP = hostIP

	entityDevice.Hostname = body.Hostname
	entityDevice.MacAddress = body.MacAddress
	// body.RemoteIP = c.
	log.Println(entityDevice)
	ct.DeviceService.Insert(entityDevice)

}

func (ct *Controller) getDevices(c *gin.Context) {
	devices := ct.DeviceService.FindAll()
	log.Println(devices)
	c.HTML(http.StatusOK, "devices.html", gin.H{
		"Devices": devices,
	})

}

func (ct *Controller) getResources(c *gin.Context) {
	resources := ct.Options.Resources
	var results []entity.Resources
	for _, resource := range resources {
		for _, ip := range resource.IP {
			var result entity.Resources
			result.ClusterName = resource.Name
			result.NodeIP = ip
			t := ct.DeviceService.GetUpdateTime(ip)
			result.DaemonUpdate = t.Format("2006-01-02 15:04:05")

			t2 := ct.FalcoService.GetUpdateTime(ip)
			result.FalcoUpdate = t2.Format("2006-01-02 15:04:05")

			results = append(results, result)
		}

	}

	log.Println(results)
	c.HTML(http.StatusOK, "resources.html", gin.H{
		"Resources": results,
	})
}

func (ct *Controller) resourcesUpdate(c *gin.Context) {
	//初始化资产
	resources := ct.Options.Resources
	var entyResource entity.Resource
	for _, resource := range resources {
		for _, ip := range resource.IP {
			entyResource.ClusterName = resource.Name
			entyResource.NodeIP = ip
			ct.ResourceService.Insert(entyResource)
		}

	}
}

func (ct *Controller) setFalco(c *gin.Context) {
	var falco entity.Falco
	if err := c.BindJSON(&falco); err != nil {
		log.Println(err)
	}
	nodeIp, err := ct.DeviceService.FindIPByAddress(falco.ClientID)
	if err != nil {
		log.Println(err)
	}
	falco.NodeIp = nodeIp
	ct.FalcoService.Insert(falco)
}

func (ct *Controller) restartFalco(c *gin.Context) {
	ip := c.Query("nodeip")
	address, _ := ct.FalcoService.FindAddressByIp(ip)
	// conn, found := ct.ClientService.GetConnection(address)
	command := "falco restart"
	ctxWithTimeout, cancel := context.WithTimeout(c, 15*time.Second)
	defer cancel()
	result, err := ct.ClientService.SendCommand(ctxWithTimeout, service.SendCommandInput{
		ClientID: address,
		Command:  command,
	})

	if err != nil {
		log.Println(err)
	}

	var commandOutput service.CommandOutput
	// json.Marshal()
	errx := json.Unmarshal(result, &commandOutput)

	if errx != nil {
		log.Println(err)
	}
	decoded := commandOutput.Response
	decode, _ := base64.StdEncoding.DecodeString(decoded)
	log.Println(string(decode))
	c.JSON(http.StatusOK, gin.H{
		"message": string(decode),
	})
	// log.Println(base64.StdEncoding.DecodeString(string(commandOutput.Response)))
}

func (ct *Controller) doCommand(c *gin.Context) {
	ip := c.Query("nodeip")
	address, _ := ct.FalcoService.FindAddressByIp(ip)
	// conn, found := ct.ClientService.GetConnection(address)
	command := "falco restart"
	ctxWithTimeout, cancel := context.WithTimeout(c, 15*time.Second)
	defer cancel()
	result, err := ct.ClientService.SendCommand(ctxWithTimeout, service.SendCommandInput{
		ClientID: address,
		Command:  command,
	})

	if err != nil {
		log.Println(err)
	}

	var commandOutput service.CommandOutput
	// json.Marshal()
	errx := json.Unmarshal(result, &commandOutput)

	if errx != nil {
		log.Println(err)
	}
	decoded := commandOutput.Response
	decode, _ := base64.StdEncoding.DecodeString(decoded)
	log.Println(string(decode))
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Restart Falco Succeed",
	// })
	// log.Println(base64.StdEncoding.DecodeString(string(commandOutput.Response)))
}

func (ct *Controller) getGenerate(c *gin.Context) {
	c.HTML(http.StatusOK, "generate.html", gin.H{
		"Address":  ct.Options.ServerAddress,
		"Port":     strings.ReplaceAll(ct.Options.ServerPort, ":", ""),
		"Filename": ct.Options.ClientName,
	})
}

func (ct *Controller) generateClient(c *gin.Context) {

	var clientGenerate entity.GenerateClient
	if err := c.ShouldBind(&clientGenerate); err != nil {
		log.Println(err)
	}
	log.Println(clientGenerate)

	var input service.BuildClientInput
	input.Filename = clientGenerate.Filename
	input.ServerAddress = clientGenerate.Address
	input.ServerPort = clientGenerate.Port
	log.Println(input)
	binary, err := ct.ClientService.BuildClient(input)
	if err != nil {
		log.Println(err)
	}
	log.Println("binaray: ", binary)
	c.String(http.StatusOK, binary)

}

// func (ct *Controller) getDevicesHandler(c *gin.Context) {
// 	devices, err := ct.DeviceService.FindAll()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.HTML(http.StatusOK, "devices.html", gin.H{
// 		"Devices": devices,
// 	})
// }

func (ct *Controller) getHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (ct *Controller) clientHandler(c *gin.Context) {
	clientID := c.GetHeader("x-client")
	token := c.GetHeader("cookie")
	log.Println(token)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error connecting client:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = ct.ClientService.AddConnection(clientID, ws)
	if err != nil {
		log.Println("error adding client:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Client connected: ", clientID)
}

func (ct *Controller) sendCommand(c *gin.Context) {
	type SendCommandRequestForm struct {
		Address   string `form:"address" binding:"required"`
		Command   string `form:"command" binding:"required"`
		Parameter string `form:"parameter"`
	}
	var form SendCommandRequestForm
	log.Println(form)
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if len(strings.TrimSpace(form.Command)) == 0 {
		c.String(http.StatusOK, "")
		return
	}
	log.Println(form.Address)

	// clientID, err := utils.DecodeBase64(form.Address)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// ctxWithTimeout, cancel := context.WithTimeout(c, 10*time.Second)
	// defer cancel()

	// output, err := h.ClientService.SendCommand(ctxWithTimeout, client.SendCommandInput{
	// 	ClientID:  clientID,
	// 	Command:   form.Command,
	// 	Parameter: form.Parameter,
	// })
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.String(http.StatusOK, output.Response)
}
