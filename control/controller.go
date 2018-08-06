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
		Db.Create(&device)
	}
}

func PutMacOnline(mac string) {
	//model.UpdateDeviceOnlineStatus(model.Database, mac, 1)
	var device model.Device = model.Device{Mac: mac}
	Db.Model(&device).Update("Online", 1)
}

func PutMacOffline(mac string) {
	//model.UpdateDeviceOnlineStatus(model.Database, mac, 0)
	var device model.Device = model.Device{Mac: mac}
	Db.Model(&device).Update("Online", 0)
}

func UpdateHeartbeat(mac, heartbeat string) {
	//model.UpdateDeviceLastHeartbeat(model.Database, mac, heartbeat)
	//TODO: move heartbeat/online to redis
	var device model.Device = model.Device{Mac: mac}
	Db.Model(&device).Update("Heartbeat", heartbeat)
}

func GetHeartbeat(mac string) int64 {
	var device model.Device = model.Device{Mac: mac}
	Db.Take(&device)
	return device.Heartbeat
}
