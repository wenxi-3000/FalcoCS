package serviceimpl

import (
	"log"
	"server/dao"
	"server/entity"
	"server/service"
	"time"
)

type deviceService struct {
	deviceDao dao.Device
}

func NewDeviceService(dD dao.Device) service.DeviceService {
	return deviceService{deviceDao: dD}
}

// func NewDeviceService(deviceDao dao.Device) service.DeviceService {
// 	var ds deviceService
// 	ds.daoDevice = deviceDao

// 	var sd service.DeviceService
// 	return sd
// }

// func (d deviceService) Insert(input entity.Device) error {
// 	log.Println(".........Insert.......")
// 	err := d.daoDevice.Insert(input)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type deviceSer struct {
// 	enti entity.Device
// }
// type Device struct {
// 	entity.DBModel
// 	Hostname   string `json:"hostname"`
// 	MacAddress string `json:"mac_address"`
// 	HostIP     string `json:"hostip"`
// }

func (d deviceService) Insert(input entity.Device) {

	device := d.deviceDao.FindByMacAddress(input.MacAddress)
	if device.MacAddress == "" {
		d.deviceDao.Insert(input)
	}
	d.deviceDao.Update(input)
	// log.Println(err)
}

func (d deviceService) FindAddressByIp(ip string) (string, error) {

	return "", nil
}

func (d deviceService) FindIPByAddress(address string) (string, error) {
	device := d.deviceDao.FindByMacAddress(address)
	return device.NodeIP, nil
}

func (d deviceService) FindAll() []entity.Device {
	devices := d.deviceDao.FindAll(time.Now().Add(time.Minute * time.Duration(-3)))
	// for index, device := range devices {
	// 	devices[index].MacAddress =
	// }
	log.Println(devices)
	return devices
}

func (d deviceService) GetUpdateTime(ip string) time.Time {
	// log.Println(ip)
	result := d.deviceDao.GetUpdateTime(ip)
	return result.UpdatedAt
}

// func FindAll() []entity.Device {
// 	device.FindAll()
// }

// type deviceService struct {
// 	Repository dao.Device
// }

// func NewDeviceService(dao dao.Device) Service {
// 	return &deviceService{Dao: dao}
// }

// func DeviceService() {
// 	devices, err := dao.FindAll(time.Now().Add(time.Minute * time.Duration(-3)))
// 	if err != nil {
// 		return nil, err
// 	}
// 	for index, device := range devices {
// 		devices[index].MacAddressBase64 = utils.EncodeBase64(device.MacAddress)
// 	}
// 	return devices, nil
// }
