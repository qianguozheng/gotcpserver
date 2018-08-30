package control

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qianguozheng/goadmin/model"
)

var Db *gorm.DB

func FindMacInDB(mac string) bool {

	var dev model.Device
	Db.Where("mac = ?", mac).First(&dev)
	return dev.Mac == mac
}

func AddMacIntoDB(mac string) {

	var device model.Device = model.Device{Mac: mac}

	if FindMacInDB(mac) == false {
		//Db.Create(&device)
		model.AddDevice(device)
	}
}

func PutMacOnline(mac string) {
	//model.UpdateDeviceOnlineStatus(model.Database, mac, 1)
	var dev model.Device
	Db.Model(&model.Device{}).Where("mac=?", mac).Find(&dev)
	Db.Debug().Model(&model.DeviceStatus{}).Where("device_refer=?", dev.Id).Update("Online", true)
}

func PutMacOffline(mac string) {
	//model.UpdateDeviceOnlineStatus(model.Database, mac, 0)
	var dev model.Device
	Db.Model(&model.Device{}).Where("mac=?", mac).Find(&dev)
	Db.Debug().Model(&model.DeviceStatus{}).Where("device_refer=?", dev.Id).Update("Online", false)
}

func UpdateHeartbeat(mac string, heartbeat int64) {
	//model.UpdateDeviceLastHeartbeat(model.Database, mac, heartbeat)
	//TODO: move heartbeat/online to redis
	var dev model.Device
	Db.Model(&model.Device{}).Where("mac=?", mac).Find(&dev)
	Db.Model(&model.DeviceStatus{}).Debug().Where("device_refer=?", dev.Id).Update("Heartbeat", heartbeat)
}

func GetHeartbeat(mac string) int64 {
	var device model.Device = model.Device{Mac: mac}
	Db.Take(&device)
	return device.Heartbeat
}
