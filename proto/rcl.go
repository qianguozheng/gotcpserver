package proto

import "github.com/qianguozheng/goadmin/model"

type RclConfig struct {
	Cmd   string        `json:"cmd"`
	SeqId string        `json:"seqId"`
	Code  string        `json:"code"`
	Data  RclConfigCore `json:"data"`
}

type RclConfigCore struct {
	Md5            string         `json:"md5"`
	Mode           int            `json:"mode"`
	Name           string         `json:"name"`
	Lan            Lan            `json:"lan"`
	Rfs            []Rfs          `json:"rfs"`
	Dhcp           Dhcp           `json:"dhcp"`
	Ssid           []model.Ssid   `json:"ssids"`
	Qos            Qos            `json:"qos"`
	TrustIps       []TrustIps     `json:"trustIps"`
	TrustDomains   TrustDomains   `json:"trustDomains"`
	HttpProxy      HttpProxy      `json:"httpProxy"`
	AutoPortalStop AutoPortalStop `json:"autoPortalStop"`
	NodeConfigUrl  string         `json:"nodeConfigUrl"`
}

type Lan struct {
	Ip   string `json:"ip"`
	Mask string `json:"mask"`
}

type Rfs struct {
	Type  string `json:"type"`
	Mode  int    `json:"mode"`
	Freq  int    `json:"frequency"`
	Power int    `json:"power"`
}

type Dhcp struct {
	StartIp   string `json:"startIp"`
	EndIp     string `json:"endIp"`
	LeaseTime int    `json:"leaseTime"`
}

//type Ssid struct {
//	Port     int    `json:"port"`
//	Name     string `json:"name"`
//	Name5g   string `json:"5gname"`
//	Url      string `json:"url"`
//	Password string `json:"password"`
//}

type Qos struct {
	UpRate   int            `json:"upRate"`
	DownRate int            `json:"downRate"`
	TcpLimit int            `json:"tcpLimit"`
	UdpLimit int            `json:"udpLimit"`
	Wans     []model.WanQos `json:"wans"`
}

//type WanQos struct {
//	Port int `json:"port"`
//	Up   int `json:"up"`
//	Down int `json:"down"`
//}

//TrustIps
type TrustIps struct {
	StartIp string `json:"startIp"`
	EndIp   string `json:"endIp"`
}

type TrustDomains []string

type HttpProxy struct {
	Enable      int    `json:"enable"`
	Regex       string `json:"regex"`
	Replacement string `json:"replacement"`
}
