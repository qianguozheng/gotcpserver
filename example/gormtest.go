package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Store device info
type Device struct {
	Mac           string `gorm: "primary_key;not null;unique"`
	Id            int64  `gorm: "AUTO_INCREMENT"`
	Name          string `gorm: "size:255"`
	Online        bool
	Heartbeat     int64
	CreatedAt     time.Time
	Mode          int `gorm:"default:1"` //Router mode AP：0， Route:1
	Version       string
	LanIp         string //Lan
	LanMask       string
	DhcpEnable    int //Dhcp
	DhcpStartIp   string
	DhcpEndIp     string
	DhcpLeaseTime int
	MultiSsid     bool   //Don't impement now
	RfType        string //2.4G
	RfMode        int
	RfFreq        int
	RfPower       int
	RfType5       string //5.8G
	RfMode5       int
	RfFreq5       int
	RfPower5      int
	Ssid          []Ssid `gorm:"foreignkey:DeviceRefer"`
	Wan           []Wan  `gorm:"foreignkey:DeviceRefer"`
	Qos           []Qos  `gorm:"foreignkey:DeviceRefer"`
	ModelType     int    `gorm:"default:0"` //AR9341,AR9344,AR9531,MT7620A,GW500...
	CloudHost     string //Cloud
	CloudPort     int    `gorm:"default:37001"`
	CloudToken    string
	Md5           string `gorm:"default:'00000000000000000000000000000000'"`
}

//Table ssid
type Ssid struct {
	Port       int
	Name       string
	Url        string
	Password   string
	DevceRefer int64
}

//Table wan
type Wan struct {
	Port          int
	Mode          int
	FixIp         string
	FixMask       string
	FixGateway    string
	PPPoEAccount  string
	PPPoEPassword string
	PrimaryDns    string
	SecondaryDns  string
	DeviceRefer   int64
}

//Table wan_qos
type WanQos struct {
	Port     int
	Up       int
	Down     int
	QosRefer int
}

//Table qos
type Qos struct {
	Id          int
	UpRate      int
	DownRate    int
	TcpLimit    int
	UdpLimit    int
	WanQos      []WanQos `gorm:"foreignkey:WanQos"`
	DeviceRefer int
}

//Store QOS
//Store Trust Domain
//Store Trust Ips
//Store CommCfg
var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "testgorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&Device{}, &Qos{}, &WanQos{}, &Wan{}, &Ssid{})

	return db
}

func GetDeviceID(mac string) int64 {
	var device Device
	DB.Where("mac=?", mac).Find(&device)
	return device.Id
}

///////////////WanQos Operation////////////////////////
func InitWanQos(qosrefer, port int) {
	def := &WanQos{
		Port:     port,
		Up:       10,
		Down:     100,
		QosRefer: qosrefer,
	}
	DB.Debug().Create(def)

}

func UpdateWanQos(wanQos WanQos) {
	DB.Debug().Model(&WanQos{}).Where("qos_refer=? and port=?", wanQos.QosRefer, wanQos.Port).Update(
		"down", wanQos.Down, "up", wanQos.Up)
}

func QueryWanQos(refer int) []WanQos {
	var wanqoss []WanQos
	DB.Debug().Find(&wanqoss, "qos_refer=?", refer)
	return wanqoss
}

func DeleteWanQos(qosrefer, port int) {
	DB.Debug().Model(&WanQos{}).Where("qos_refer=? and port=?", qosrefer, port).Delete(WanQos{})
}

func TestWanQos() {

	//Add
	InitWanQos(12, 0)
	InitWanQos(12, 1)
	InitWanQos(12, 2)
	InitWanQos(12, 3)
	InitWanQos(12, 4)

	//Retrive
	x := QueryWanQos(12)

	for k, v := range x {
		fmt.Println("k,v", k, v)
	}

	del := WanQos{
		Port:     3,
		QosRefer: 12,
	}
	//Update
	del.Port = 2
	del.Down = 40
	del.Up = 5
	UpdateWanQos(del)

	x = QueryWanQos(12)

	for k, v := range x {
		fmt.Println(k, v)
	}
	//Delete
	DeleteWanQos(12, 2)

	x = QueryWanQos(12)

	for k, v := range x {
		fmt.Println(k, v)
	}

}

///////////////Qos Operation//////////////////////////
func InitQos(refer int) {
	qos := &Qos{
		UpRate:      200,
		DownRate:    4096,
		TcpLimit:    200,
		UdpLimit:    100,
		DeviceRefer: refer,
	}
	if false == DB.Debug().NewRecord(qos) {
		DB.Debug().Create(qos)
	}

}
func UpdateQos(qos Qos) {
	DB.Debug().Model(&Qos{}).Where("device_refer=?", qos.DeviceRefer).Update(&qos)
}

func QueryQos(refer int) Qos {
	var qos Qos
	DB.Debug().Find(&qos, "device_refer=?", refer)
	return qos
}

func TestQos() {
	InitQos(13)
	qos := Qos{
		DeviceRefer: 13,
		UpRate:      90,
		DownRate:    399,
	}

	q := QueryQos(13)
	fmt.Println(q)

	UpdateQos(qos)

	q = QueryQos(13)
	fmt.Println(q)
}

func main() {

	DB = InitDB()
	defer DB.Close()

	//TestWanQos()
	TestQos()
}
