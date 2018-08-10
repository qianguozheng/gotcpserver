package control

import (
	"encoding/json"
	"fmt"

	log "../log"
	"../proto"
	"github.com/qianguozheng/goadmin/model"
)

//Retrive Device Information to fill RCL config structure

func RetriveDeviceInfoFormRCL(mac string) string {

	dev := model.GetDeviceByMac(mac)
	fmt.Println(dev)
	ssid := model.GetSsidByDeviceId(dev.Id)
	qos := model.GetQosByDeviceId(dev.Id)
	wanqos := model.GetWanQosByQosId(qos.Id)
	md5 := model.GetMd5ByDeviceId(dev.Id)

	if 0 == dev.Sync {
		md5.Md5 = "00000000000000000000000000000000"
	}

	rclCore := proto.RclConfigCore{
		Md5:  md5.Md5,
		Mode: dev.Mode,
		Name: dev.Name,
		Lan: proto.Lan{
			Ip:   dev.LanIp,
			Mask: dev.LanMask,
		},
		Rfs: []proto.Rfs{
			{
				Type:  dev.RfType,
				Mode:  dev.RfMode,
				Freq:  dev.RfFreq,
				Power: dev.RfPower,
			},
			{
				Type:  dev.RfType5,
				Mode:  dev.RfMode5,
				Freq:  dev.RfFreq5,
				Power: dev.RfPower5,
			},
		},
		Dhcp: proto.Dhcp{
			StartIp:   dev.DhcpStartIp,
			EndIp:     dev.DhcpEndIp,
			LeaseTime: dev.DhcpLeaseTime,
		},
		Ssid: ssid,
		Qos: proto.Qos{
			UpRate:   qos.UpRate,
			DownRate: qos.DownRate,
			TcpLimit: qos.TcpLimit,
			UdpLimit: qos.UdpLimit,
			Wans:     wanqos,
		},
		TrustIps:      proto.TrustIps{},
		TrustDomains:  proto.TrustDomains{},
		HttpProxy:     proto.HttpProxy{},
		NodeConfigUrl: "http://cdn.magicwifi.com.cn/idc/nodes.json",
		AutoPortalStop: proto.AutoPortalStop{
			IOSEnable:     0,
			AndroidEnable: 0,
			Ios: proto.ReqHdr{
				Host:      []string{},
				Uri:       []string{},
				UserAgent: []string{"captivenetworksupport"},
			},
			Android: proto.ReqHdr{
				Host:      []string{"vivo.com.cn", "clients3.google.com"},
				Uri:       []string{"/generate_204", "/rsp204", "/qbprobe/netprobe.txt"},
				UserAgent: []string{},
			},
		},
	}

	rcl := proto.RclConfig{
		Cmd:   "rcl",
		SeqId: "uniqueid",
		Code:  "000",
		Data:  rclCore,
	}
	//TODO: checkValid
	//Type - GW500/AR9344 have different requirement
	//Mode - AP/Route have different requirement

	rclByte, err := json.Marshal(rcl)
	if err != nil {
		log.Error("marshal json failed")
		return ""
	}
	log.Debug("rcl:", string(rclByte))
	return string(rclByte)
}
