package control

import (
	"encoding/json"

	log "../log"
	"../proto"
	"github.com/qianguozheng/goadmin/model"
)

//type ConfigReadCore struct {
//	Wans []WanConfig  `json:"wans"`
//	Rfs  []Rfs        `json:rfs`
//	Ssid []model.Ssid `json:ssids`
//	Qos  Qos          `json:qos`
//}

func ParseConfigReadMsg(msg []byte, mac string) {
	var configRead proto.ConfigRead

	if err := json.Unmarshal(msg, &configRead); err == nil {
		log.Debug("ParseConfigRead:%v", configRead)
		configRead := configRead.Data
		//TODO: store data to database

		dev := model.GetDeviceByMac(mac)
		dev.Mode = configRead.Mode
		dev.CloudHost = configRead.CC.Host
		dev.CloudPort = configRead.CC.Port
		dev.CloudToken = configRead.CC.Token
		dev.LanIp = configRead.Lan.Ip
		dev.LanMask = configRead.Lan.Mask
		dev.DhcpStartIp = configRead.Dhcp.StartIp
		dev.DhcpEndIp = configRead.Dhcp.EndIp
		dev.DhcpLeaseTime = configRead.Dhcp.LeaseTime

		dev.Sync = 1 //mark read from device

		//rfs
		for _, v := range configRead.Rfs {
			if v.Type == "2.4G" {
				dev.RfType = "2.4G"
				dev.RfFreq = v.Freq
				dev.RfMode = v.Mode
				dev.RfPower = v.Power
			} else if v.Type == "5G" {
				dev.RfType5 = "5G"
				dev.RfFreq5 = v.Freq
				dev.RfMode5 = v.Mode
				dev.RfPower5 = v.Power
			}
		}
		model.UpdateDevice(dev)

		//wans
		model.DeleteWanByDeviceId(dev.Id)
		for _, wan := range configRead.Wans {
			wan.DeviceRefer = dev.Id
			//			fmt.Println("add wan to db", wan)
			model.AddWan(wan)
		}
		//Ssid
		for k, _ := range configRead.Ssid {
			configRead.Ssid[k].DeviceRefer = dev.Id
		}

		model.AddSsid(configRead.Ssid)
		//Qos
		qos := model.Qos{
			UpRate:      configRead.Qos.UpRate,
			DownRate:    configRead.Qos.DownRate,
			TcpLimit:    configRead.Qos.TcpLimit,
			UdpLimit:    configRead.Qos.UdpLimit,
			DeviceRefer: dev.Id,
		}
		model.UpdateQos(qos)
		//wanqos
		qos = model.GetQosByDeviceId(dev.Id)
		model.DeleteWanQosByQosId(qos.Id)
		for _, v := range configRead.Qos.Wans {
			v.QosRefer = qos.Id
			model.AddWanQosConfig(v)
		}

	}
}
