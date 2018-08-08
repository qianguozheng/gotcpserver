package server

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"../control"

	"../proto"
)

//TODO： implement the minimal of firmware
func processCommand(dat map[string]interface{}, conn net.Conn) (string, uint32) {

	//TODO: make verification work
	if dat["cmd"] == "verification" {
		vdata := proto.VerificationData{
			TerminalMac: dat["terminalMac"].(string),
			Valid:       1441,
			AuthType:    0,
			AuthId:      "13538273761",
			UpRate:      100,
			DownRate:    101,
			TcpLimit:    222,
			UdpLimit:    223,
		}
		verify := proto.Verification{
			Cmd:   dat["cmd"].(string),
			SeqId: dat["seqId"].(string),
			Code:  "000",
			Data:  vdata,
		}
		result, err := json.Marshal(verify)
		if err == nil {
			fmt.Println("verification==", string(result))
			return string(result), proto.CmdKV[dat["cmd"].(string)]
		}
	}

	//TODO: implement rcl config
	if dat["cmd"].(string) == "rcl" {
		fmt.Println("Rcl request")

		result := control.RetriveDeviceInfoFormRCL(dat["mac"].(string))
		return result, proto.CmdKV[dat["cmd"].(string)]
		//MTK
		//return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.8.200\",\"leaseTime\":120,\"startIp\":\"192.168.8.100\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.8.1\",\"mask\":\"255.255.255.0\"},\"md5\":\"79a8836d586f60d998662169b4ff0f48\",\"mode\":1,\"name\":\"-MagicWiFi-58D1-MTK\",\"qos\":{\"downRate\":1024,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":10240,\"port\":0,\"up\":2048}]},\"rfs\":[{\"frequency\":0,\"mode\":0,\"power\":100,\"type\":\"2.4\"}],\"ssids\":[{\"name\":\"-MagicWiFi-_58D0\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/captive/index.html\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"},{\"endIp\":\"  192.168.70.43\",\"startIp\":\"192.168.70.43  \"},{\"endIp\":\"183.238.95.110\",\"startIp\":\"183.238.95.101\"}],\"autoPortalStop\":{\"iOSEnable\":1,\"androidEnable\":1,\"ios\":{\"userAgent\":[\"captivenetworksupport\"]},\"android\":{\"uri\":[\"/generate_204\"],\"host\":[\"vivo.com.cn\"]}},\"auditEnable\":1},\"seqId\":\"87e41fcbfd95430e8f05f0f8ebf53273\"}"), proto.CmdKV[dat["cmd"].(string)]

		//x86
		return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.60.127\",\"leaseTime\":120,\"startIp\":\"192.168.48.200\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.48.1\",\"mask\":\"255.255.240.0\"},\"md5\":\"30d0a580df0e4f3826ce1999f3aa5469\",\"mode\":1,\"name\":\"Richard-X86\",\"qos\":{\"downRate\":512,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":0,\"port\":0,\"up\":0},{\"down\":0,\"port\":1,\"up\":0},{\"down\":0,\"port\":2,\"up\":0},{\"down\":0,\"port\":3,\"up\":0},{\"down\":0,\"port\":4,\"up\":0}]},\"ssids\":[{\"name\":\"-MagicWiFi-invalid\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/captive/index.html\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"}],\"autoPortalStop\":{\"iOSEnable\":1,\"androidEnable\":1,\"ios\":{\"userAgent\":[\"captivenetworksupport\",\"helloworld\",\"applexxxx\"],\"uri\":[\"/genrate_204\",\"/qbprobenet.txt\",\"/checknetwork\"],\"host\":[\"test.com\"]},\"android\":{\"uri\":[\"/generate_204\",\"/genrate_2004\",\"/qbprobenet.txt\"],\"host\":[\"vivo.com.cn\",\"clients3.google.com\",\"holyshit.com\"],\"userAgent\":[\"androidxxxx\",\"fuckyou\"]}}}}"), proto.CmdKV[dat["cmd"].(string)]
		//return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.60.127\",\"leaseTime\":120,\"startIp\":\"192.168.48.200\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.48.1\",\"mask\":\"d255.255.240.0\"},\"md5\":\"70a8836d586f60d998562169b4ff0f48\",\"mode\":1,\"name\":\"Richard-X86\",\"qos\":{\"downRate\":512,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":0,\"port\":0,\"up\":0},{\"down\":0,\"port\":1,\"up\":0},{\"down\":0,\"port\":2,\"up\":0},{\"down\":0,\"port\":3,\"up\":0},{\"down\":0,\"port\":4,\"up\":0}]},\"ssids\":[{\"name\":\"-MagicWiFi-invalid\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/captive/index.html\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"}],\"auditEnable\":1,\"autoPortalStop\":{\"iOSEnable\":0,\"androidEnable\":0,\"ios\":{\"userAgent\":[\"captivenetworksupport\"]},\"android\":{\"uri\":[\"/generate_204\",\"/qbprobenet.txt\"],\"host\":[\"vivo.com.cn\",\"clients3.google.com\"]}}}}"), proto.CmdKV[dat["cmd"].(string)]

		//AR9344
		//return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.8.200\",\"leaseTime\":120,\"startIp\":\"192.168.8.100\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.8.1\",\"mask\":\"255.255.255.0\"},\"md5\":\"70a8836d586f60d998562169b4ff0f49\",\"mode\":1,\"name\":\"-MagicWiFi-518\",\"qos\":{\"downRate\":1024,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":10240,\"port\":0,\"up\":2048}]},\"rfs\":[{\"frequency\":10,\"mode\":0,\"power\":30,\"type\":\"2.4\"}],\"ssids\":[{\"name\":\"-MagicWiFi-518\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/portal/\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"},{\"endIp\":\"  192.168.70.43\",\"startIp\":\"192.168.70.43  \"},{\"endIp\":\"183.238.95.110\",\"startIp\":\"183.238.95.101\"}],\"autoPortalStop\":{\"iOSEnable\":1,\"androidEnable\":1,\"ios\":{\"userAgent\":[\"captivenetworksupport\",\"helloworld\",\"applexxxx\"],\"uri\":[\"/genrate_204\",\"/qbprobenet.txt\",\"/checknetwork\"],\"host\":[\"test.com\"]},\"android\":{\"uri\":[\"/generate_204\",\"/genrate_2004\",\"/qbprobenet.txt\"],\"host\":[\"vivo.com.cn\",\"clients3.google.com\",\"holyshit.com\"],\"userAgent\":[\"androidxxxx\",\"fuckyou\"]}}},\"seqId\":\"306ff4d0c3a94923a646df3d53d7dbac\"}"),proto.CmdKV[dat["cmd"].(string)]
		//return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.8.200\",\"leaseTime\":120,\"startIp\":\"192.168.8.100\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.8.1\",\"mask\":\"255.255.255.0\"},\"md5\":\"70a8836d586f60d998562169b4ff0f49\",\"mode\":1,\"name\":\"-MagicWiFi-518\",\"qos\":{\"downRate\":1024,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":10240,\"port\":0,\"up\":2048}]},\"rfs\":[{\"frequency\":10,\"mode\":0,\"power\":30,\"type\":\"2.4\"}],\"ssids\":[{\"name\":\"-MagicWiFi-518\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/portal/\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"},{\"endIp\":\"  192.168.70.43\",\"startIp\":\"192.168.70.43  \"},{\"endIp\":\"183.238.95.110\",\"startIp\":\"183.238.95.101\"}],\"autoPortalStop\":{\"iOSEnable\":0,\"androidEnable\":0,\"ios\":{\"userAgent\":[\"captivenetworksupport\"]},\"android\":{\"uri\":[\"/generate_204\",\"/qbprobe/netprobe.txt\"],\"host\":[\"vivo.com.cn\",\"clients3.google.com\"]}}},\"seqId\":\"306ff4d0c3a94923a646df3d53d7dbac\"}"),proto.CmdKV[dat["cmd"].(string)]

		//AR9341
		//return string("{\"cmd\":\"rcl\",\"code\":\"000\",\"data\":{\"dhcp\":{\"endIp\":\"192.168.8.200\",\"leaseTime\":120,\"startIp\":\"192.168.8.100\"},\"httpProxy\":{\"enable\":false,\"regex\":\"\",\"replacement\":\"\"},\"lan\":{\"ip\":\"192.168.8.1\",\"mask\":\"255.255.255.0\"},\"md5\":\"442d695ae1dce3a1f7e2974f9b110ff0\",\"mode\":1,\"name\":\"QA-TEST-AP-9341\",\"qos\":{\"downRate\":1024,\"tcpLimit\":200,\"udpLimit\":60,\"upRate\":128,\"wans\":[{\"down\":10240,\"port\":0,\"up\":2048}]},\"rfs\":[{\"frequency\":11,\"mode\":0,\"power\":20,\"type\":\"2.4\"}],\"ssids\":[{\"name\":\"-MagicWiFi-QA-9341\",\"password\":\"\",\"port\":0,\"url\":\"http://magicwifi.com.cn/portal/\"}],\"trustDomains\":[\"alicdn.com\",\"alipaydns.com\",\"hi-wifi.cn\",\"sasdk.cs0309.3g.qq.com\",\"wx.tenpay.com\",\"tbcache.com\",\"api.apkplug.com\",\"alog.umeng.com\",\"alipay.com\",\"pay.heepay.com\",\"wx.tencent.cn\",\"magicwifi.com.cn\",\"sasdk.3g.qq.com\",\"mp.weixin.qq.com\",\"zhongxin.junka.com\",\"rippletek.com\",\"alipayobjects.com\"],\"trustIps\":[{\"endIp\":\"183.232.2.4\",\"startIp\":\"183.232.2.4\"},{\"endIp\":\"121.201.55.228\",\"startIp\":\"121.201.55.153\"},{\"endIp\":\"123.206.2.126\",\"startIp\":\"123.206.2.126\"},{\"endIp\":\"123.206.2.188\",\"startIp\":\"123.206.2.188\"}],\"autoPortalStop\":{\"iOSEnable\":0,\"androidEnable\":0,\"ios\":{\"userAgent\":[\"captivenetworksupport\",\"helloworld\",\"applexxxx\"],\"uri\":[\"/genrate_204\",\"/qbprobenet.txt\",\"/checknetwork\"],\"host\":[\"test.com\"]},\"android\":{\"uri\":[\"/generate_204\",\"/genrate_2004\",\"/qbprobenet.txt\"],\"host\":[\"vivo.com.cn\",\"clients3.google.com\",\"holyshit.com\"],\"userAgent\":[\"androidxxxx\",\"fuckyou\"]}}},\"seqId\":\"69a629b9e1ae4a1a933e0d4e851040d5\"}"), proto.CmdKV[dat["cmd"].(string)]
	}

	if dat["cmd"].(string) == "login" {
		fmt.Println("process login")

		i := dat["mac"]
		if i == nil {
			return "", proto.CmdKV[dat["cmd"].(string)]
		}
		mac := i.(string)
		/// Save conn into ConnMap for later usage
		ConnMap[mac] = conn

		fmt.Println("login mac:", mac)
		found := control.FindMacInDB(mac)
		fmt.Println("found", found)
		if false == found {
			fmt.Println("add mac info db")
			control.AddMacIntoDB(mac)
			control.PutMacOnline(mac)
		}

		login := proto.RespParam{
			Cmd:   dat["cmd"].(string),
			SeqId: dat["seqId"].(string),
			Code:  "000",
			Data:  []string{"login"},
		}
		result, err := json.Marshal(login)
		if err == nil {
			fmt.Println("login request=", string(result))
			return string(result), proto.CmdKV[dat["cmd"].(string)]
		}

	}

	if dat["cmd"].(string) == "heartbeat" {
		mac := dat["mac"].(string)
		control.UpdateHeartbeat(mac, time.Now().Format("2006-01-02 15:04:05"))
		heartbeat := proto.RespParam{
			Cmd:   dat["cmd"].(string),
			SeqId: dat["seqId"].(string),
			Code:  "000",
			Data:  []string{"test"},
		}
		result, err := json.Marshal(heartbeat)
		if err == nil {
			fmt.Println("[ToClient]=", string(result))
			fmt.Println("=============================================")
			return string(result), proto.CmdKV[dat["cmd"].(string)]
		}
	}

	if dat["cmd"].(string) != "" {
		if dat["cmd"].(string) == "web_read_resp" ||
			dat["cmd"].(string) == "resource_read_resp" ||
			dat["cmd"].(string) == "web_write_resp" ||
			dat["cmd"].(string) == "resource_write_resp" {
			//return "",0
		}
		login := proto.RespParam{
			Cmd:   dat["cmd"].(string),
			SeqId: dat["seqId"].(string),
			Code:  "000",
			Data:  []string{"test"},
		}
		result, err := json.Marshal(login)
		if err == nil {
			fmt.Println("[ToClient]=", string(result))
			fmt.Println("=============================================")
			return string(result), proto.CmdKV[dat["cmd"].(string)]
		}
		//fmt.Println("result=", err.Error())
	}

	return "{\"cmd\":\"not found cmd\"}", 0
}

func FindMacInConnMap(conn net.Conn) string {
	for k, v := range ConnMap {
		fmt.Println("k=", k, " v=", v)
		if v == conn {
			return k
		}
	}
	return ""
}
func handleMsg(msg []byte, conn net.Conn) (string, uint32) {

	dat := make(map[string]interface{})
	fmt.Println("handleMsg..")

	if err := json.Unmarshal(msg, &dat); err == nil {

		cmd := dat["cmd"].(string)
		//TODO: response from client
		if cmd == "reboot_resp" ||
			cmd == "upgrade_resp" ||
			cmd == "cc_write_resp" ||
			cmd == "rc_write_resp" ||
			cmd == "upgrade_resp" ||
			cmd == "config_read_resp" ||
			cmd == "notification_resp" {

			if cmd == "config_read_resp" {
				//Find mac by net.Conn
				mac := FindMacInConnMap(conn)
				control.ParseConfigReadMsg(msg, mac)
			}

			return string(msg), proto.CmdKV[dat["cmd"].(string)]
		}
		//TODO: send msg to rpc client for Authentication
		if cmd == "verification_req" {

			resp, err := RPCClientRequest(string(msg))
			if err != nil {
				return "", 0
			} else {
				return resp, proto.CmdKV["verification_resp"]
			}
		}

		return processCommand(dat, conn)

	}
	return "{\"CMD\":\"Invalid Json Format\"}", 0
}
