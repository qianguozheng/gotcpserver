package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Store device info
type Device struct {
	Mac           string `gorm: "primary_key;not null;unique"`
	Id            int64  `gorm: "AUTO_INCREMENT"`
	Name          string `gorm: "size:255"`
	Online        bool
	Heartbeat     string
	CreatedAt     time.Time
	Mode          int `gorm:"default:1"` //Router mode AP：0， Route:1
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
	MultiSsid     bool   //Don't impement now
	RfType        string //2.4G
	RfMode        int
	RfFreq        int
	RfPower       int
	RfType5       string //5.8G
	RfMode5       int
	RfFreq5       int
	RfPower5      int

	ModelType  int    `gorm:"default:0"` //AR9341,AR9344,AR9531,MT7620A,GW500...
	CloudHost  string //Cloud
	CloudPort  int    `gorm:"default:37001"`
	CloudToken string
	Md5        string `gorm:"default:'00000000000000000000000000000000'"`
}
