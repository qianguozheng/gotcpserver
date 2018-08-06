package main

import (
	"fmt"

	"github.com/valyala/gorpc"
)

func DispatcherAddFunc() *gorpc.Dispatcher {
	d := gorpc.NewDispatcher()

	d.AddFunc("SendCommand", func(request string) (string, error) {
		//Parse Json
		//return ExecSendCommand(request)
		return request, nil
	})

	//TODO: Add RPC FUNC here,

	return d

}

func RPCClientRequest() {
	d := DispatcherAddFunc()
	//Client
	c := gorpc.NewTCPClient("127.0.0.1:12445")
	c.Start()
	defer c.Stop()

	dc := d.NewFuncClient(c)

	res, err := dc.Call("SendCommand", "{\"cmd\":\"reboot\", \"mac\":\"00782fe82e35\",\"seqId\":\"1243324235\"}")
	fmt.Printf("res:%s, err:%v\n", res, err)
	//TODO: Process Response
}

func RPCServerService() *gorpc.Server {
	d := DispatcherAddFunc()

	//Start TCP Server
	s := gorpc.NewTCPServer("127.0.0.1:12445", d.NewHandlerFunc())
	if err := s.Start(); err != nil {
		fmt.Println("Cannot start rpc server")
	}
	//defer s.Stop()
	return s
}

func main() {
	//	s := RPCServerService()
	//	defer s.Stop()

	RPCClientRequest()

}
