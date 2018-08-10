package main

import (
	"fmt"
	"net"
)

func GetInterfaceAddr(iface string) string {
	var ifi *net.Interface
	var err error
	if ifi, err = net.InterfaceByName(iface); err != nil {
		fmt.Println("Interface Error:", err.Error())
		return ""
	}

	addrs, err := ifi.Addrs()
	checkError(err, "addrs:")

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	addr := GetInterfaceAddr("enp2s0")
	fmt.Println("Addr:", addr)
	GetInterfaceAddr("lo")
	GetInterfaceAddr("docker0")
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error())
	}
}
