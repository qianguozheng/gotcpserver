package proto

import (
	"github.com/qianguozheng/goadmin/model"
)

type ConfigRead struct {
	Cmd   string         `json:"cmd"`
	SeqId string         `json:"seqId"`
	Code  string         `json:"code"`
	Data  ConfigReadCore `json:"data"`
}
type ConfigReadCore struct {
	Mode int          `json:"mode"`
	CC   CloudConfig  `json:"cc"`
	Wans []model.Wan  `json:"wans"`
	Lan  Lan          `json:"lan"`
	Rfs  []Rfs        `json:"rfs"`
	Dhcp Dhcp         `json:"dhcp"`
	Ssid []model.Ssid `json:"ssids"`
	Qos  Qos          `json:"qos"`
}

type CloudConfig struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}
