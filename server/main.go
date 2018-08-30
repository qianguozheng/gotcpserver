package server

import (
	"net"

	"../control"
	"../log"
	"../proto"
	"github.com/qianguozheng/goadmin/model"
)

//http://www.01happy.com/golang-tcp-socket-adhere/

func Main() {
	//TODO: Dynamic assign ip and port
	Opts = parseArgs()

	log.LogTo(Opts.LogTo, Opts.LogLevel)

	var hostport string
	if Opts.TcpAddr == "" || Opts.Port == "" {
		log.Error("tcp and port MUST not be null")
		return
	}
	//HostAndPort := "192.168.0.12:37001"
	hostport = Opts.TcpAddr + Opts.Port
	listener := server(hostport)
	defer listener.Close()

	CommInit()

	log.Info("DBPath: %s, DBType: %s", Opts.DBPath, Opts.DBType)
	//control.Db = model.InitDB(Opts.DBPath)
	control.Db = model.InitDB(Opts.DBPath, Opts.DBType)
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

	log.Info("Listening to: %s", listener.Addr().String())
	return listener
}
