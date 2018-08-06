package server

import (
	"net"

	"../control"
	"../proto"
	"github.com/qianguozheng/goadmin/model"
)

//http://www.01happy.com/golang-tcp-socket-adhere/

const (
	MaxRead = 1024 * 1024 //1MB
)

//Store net.Conn of client into map, so we can send command to client via net.Conn
var ConnMap map[string]net.Conn
var RpcResponse chan interface{}

func GobalInit() {
	RpcResponse = make(chan interface{}, 1)
	ConnMap = make(map[string]net.Conn)
}
func Main() {
	//TODO: Dynamic assign ip and port
	HostAndPort := "192.168.0.12:37001"
	listener := server(HostAndPort)
	defer listener.Close()

	GobalInit()

	control.Db = model.InitDB()
	model.DB = control.Db
	defer control.Db.Close()

	s := RPCServerService()
	defer s.Stop()

	proto.InitKeyValue()

	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")

		go connectionHandler(conn)
	}
}

//Listen to socket and return *net.TCPListener
func server(host string) *net.TCPListener {

	server, err := net.ResolveTCPAddr("tcp", host)
	checkError(err, "Resolving address:port failed: "+host)

	listener, err := net.ListenTCP("tcp", server)
	checkError(err, "ListenTCP: ")

	println("Listening to: ", listener.Addr().String())
	return listener
}
