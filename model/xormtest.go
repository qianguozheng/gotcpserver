package model

/*
import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'usr_name'"`
}

//Store device info
type Device struct {
	Id            int64  `xorm: "autoincr"`
	Name          string `xorm: "varchar(64)"`
	Mac           string `xorm: "pk varchar(12) notnull unique"`
	Online        bool
	Heartbeat     int64
	CreatedAt     time.Time `xorm:"created"`
	Mode          int       //Router mode AP：0， Route:1
	Version       string
	LanIp         string //Lan
	LanMask       string
	DhcpEnable    int //Dhcp
	DhcpStartIp   string
	DhcpEndIp     string
	DhcpLeaseTime int
	SsidPort      int //SSID
	SsidName      string
	SsidUrl       string
	SsidPassword  string
	MultiSsid     bool

	CloudHost  string //Cloud
	CloudPort  int
	CloudToken string
}

//Store LAN
type Lan struct {
	Ip   string
	Mask string
}

//Store DHCP
type Dhcp struct {
	Enable    int
	StartIp   string
	EndIp     string
	LeaseTime int
}

//Store SSID
type Ssid struct {
	Port     int
	Name     string
	Url      string
	Password string
}

//Store RF, dual band
type Rf struct {
	Type  string
	Mode  int
	Freq  int
	Power int
}

//Store Cloud Info
type Cloud struct {
	Host  string
	Port  int
	Token string
}

//Store QOS
//Store Trust Domain
//Store Trust Ips
//Store CommCfg

func main() {
	orm, err := xorm.NewEngine("sqlite3", "./xorm.db")
	err = orm.Sync(new(User), new(Device))
	//defer
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL(true)

	testDevice(orm)
	orm.Close()
}

func testDevice(orm *xorm.Engine) {
	//Insert

	device := new(Device)
	device.Mac = "123456789012"
	device.LanIp = "192.168.1.1"
	device.LanMask = "255.255.255.0"
	affected, err := orm.Insert(device)
	fmt.Println("devid, affected:", device.Id, affected, err)

	//Batch Insert
	devs := make([]Device, 3)
	devs[0].Mac = "helloworld"
	devs[1].Mac = "test...."
	devs[2].Mac = "xxxxttttt"
	affected, err = orm.Insert(&devs)
	fmt.Println("affected:", affected)

	//Delete by Id
	dd := new(Device)
	affected, err = orm.Id(5).Delete(dd)
	ddd := &Device{Mac: "helloworld"}
	affected, err = orm.Delete(ddd)

	//Query
	dev2 := &Device{Mac: "123456789012"}
	has, err := orm.Get(dev2)

	fmt.Println("has:", has, dev2)
	devices := make([]Device, 0)

	//Query
	orm.Sql("select * from device").Find(&devices)
	for _, d := range devices {
		fmt.Println(d.Id, d.Mac, d.CreatedAt)
	}

	//	results, err := orm.Query("select * from device")
	//	//	fmt.Println("query result:", results)
	//	for _, x := range results {
	//		for mm, t := range x {
	//			fmt.Println(mm, t)
	//		}
	//	}
	//	//fmt.Println(orm)
}
*/
