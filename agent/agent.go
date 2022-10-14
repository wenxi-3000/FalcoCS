package main

import (
	"FalcoCS/agent/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PostJsonData struct {
	HostName string              `json:"hostname"`
	HostIp   map[string]struct{} `json:"hostip"`
}

func New() []byte {
	var postJsonData PostJsonData
	collection := collection.NewAgentInfo()
	postJsonData.HostName = collection.HostName
	postJsonData.HostIp = collection.HostIp
	result, err := json.Marshal(postJsonData)
	if err != nil {
		log.Println(err)
	}
	return result
}

func main() {
	postJsonData := New()
	Sender(postJsonData)

}

func Sender(data []byte) {
	// data := make(map[string]interface{})
	// data["HostName"] = collection.HostName
	// data["HostIp"] = collection.HostIp
	// byteData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Println(err)
	// }
	req := bytes.NewBuffer([]byte(data))

	url := "http://172.16.42.150:8081/devices"
	request, err := http.NewRequest("POST", url, req)
	if err != nil {
		log.Println(err)
	}
	log.Println(request)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(resp)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(respBytes)
	}

	fmt.Println(string(respBytes))

}
