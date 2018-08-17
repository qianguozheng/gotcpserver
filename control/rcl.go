package control

import (
	"encoding/json"
	"fmt"
	"strings"

	log "../log"
	"../proto"
	"github.com/qianguozheng/goadmin/model"
)

//Retrive Device Information to fill RCL config structure

func RetriveDeviceInfoFormRCL(mac string, rclEnable bool) string {

	dev := model.GetDeviceByMac(mac)
	fmt.Println(dev)
	ssid := model.GetSsidByDeviceId(dev.Id)
	qos := model.GetQosByDeviceId(dev.Id)
	wanqos := model.GetWanQosByQosId(qos.Id)
	md5 := model.GetMd5ByDeviceId(dev.Id)

	//TrustDomains
	domains := model.GetDomainsByGloabl()
	domains2 := model.GetDomainsByProjectId(dev.ProjectRefer)
	domains = append(domains, domains2...)
	domain := proto.TrustDomains{}

	for _, v := range domains {
		domain = append(domain, v.Domain)
	}

	//TrustIps
	ips := model.GetIpsByProjectId(dev.ProjectRefer)

	//	for _, v := range ips {
	//		fmt.Println("v=", v.Ip)
	//	}
	ips2 := model.GetIpsByGlobal()
	for _, v := range ips2 {
		found := false
		for _, vv := range ips {
			if v.Ip == vv.Ip {
				found = true
			}
		}
		if !found {
			ips = append(ips, v)
		}
	}
	//ips = append(ips, ips2...)

	trustIps := []proto.TrustIps{}
	for _, v := range ips {
		var trustIp proto.TrustIps
		if strings.Contains(v.Ip, "-") {
			vv := strings.Split(v.Ip, "-")
			//for _, ip := range vv {
			trustIp.StartIp = vv[0]
			trustIp.EndIp = vv[1]
			//}
		} else {
			trustIp.StartIp = v.Ip
			trustIp.EndIp = v.Ip
		}
		trustIps = append(trustIps, trustIp)
	}

	//DnsBogus
	dnsBogus := []proto.DnsBogus{}
	bogus := model.GetDnsBogusByProjectId(dev.ProjectRefer)
	for _, v := range bogus {
		if v.Status == 2 {
			dns := proto.DnsBogus{
				Domain: v.Domain,
				Host:   v.Ip,
			}
			dnsBogus = append(dnsBogus, dns)
		}
	}

	if 0 == dev.Sync || !rclEnable {
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
		TrustIps:      trustIps,
		TrustDomains:  domain,
		DnsBogus:      dnsBogus,
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
	log.Debug("rcl: %s", string(rclByte))
	return string(rclByte)
}
