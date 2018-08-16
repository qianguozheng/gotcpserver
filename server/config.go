package server

import (
	"flag"
	"fmt"
	"net"

	log "../log"
)

const (
	MaxRead = 1024 * 1024 //1MB
)

var Opts *Options

type Options struct {
	TcpAddr  string
	Port     string
	LogTo    string
	LogLevel string
	DBPath   string
	Rcl      bool
}

func parseArgs() *Options {
	tcpAddr := flag.String("addr", "112.74.112.103", "address to serve the tcp service")
	portPtr := flag.String("port", ":80", "port to serve the service")
	dbPath := flag.String("dbpath", "../goadmin/goadmin.db", "path to sqlite3 database")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	rcl := flag.Bool("rcl", false, "Enable Remote control local support, default: false")

	flag.Parse()

	return &Options{
		TcpAddr:  *tcpAddr,
		Port:     *portPtr,
		LogTo:    *logto,
		LogLevel: *loglevel,
		DBPath:   *dbPath,
		Rcl:      *rcl,
	}
}

func GetInterfaceAddr(iface string) string {
	var ifi *net.Interface
	var err error
	if ifi, err = net.InterfaceByName(iface); err != nil {
		log.Error("Interface Error:", err.Error())
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
