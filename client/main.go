package main

import (
	"bytes"
	"client/device"
	"client/falco"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	device := device.NewDevice()
	falco := falco.NewFalco()
	go doSenderDevice(device)
	go doSenderFalco(falco)
	for {
		time.Sleep(10 * time.Second)
	}
}

func doSenderFalco(falco []byte) {
	for {
		SenderFalco(falco)
		time.Sleep(3 * time.Second)
	}
}

func SenderFalco(falco []byte) {
	data := bytes.NewBuffer([]byte(falco))

	url := "http://172.16.42.150:8081/falco"
	request, err := http.NewRequest("POST", url, data)
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

func doSenderDevice(device []byte) {
	for {
		SenderDevice(device)
		time.Sleep(30 * time.Second)
	}

}

func SenderDevice(device []byte) {
	// data := make(map[string]interface{})
	// data["HostName"] = collection.HostName
	// data["HostIp"] = collection.HostIp
	// byteData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Println(err)s
	// }
	data := bytes.NewBuffer([]byte(device))

	url := "http://172.16.42.150:8081/device"
	request, err := http.NewRequest("POST", url, data)
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
