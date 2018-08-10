package server

import (
	"net"
	"strings"
)

/*
	Internal Communication Structure
*/

type CommUnit struct {
	ConnMap map[string]*net.Conn
	RpcCh   map[string]chan interface{}
}

var Comm *CommUnit

func NewCommUnit() *CommUnit {
	comm := &CommUnit{
		ConnMap: make(map[string]*net.Conn),
		RpcCh:   make(map[string]chan interface{}),
	}
	return comm
}

func (c *CommUnit) AddConn(mac string, conn *net.Conn) {
	c.ConnMap[mac] = conn
}

func (c *CommUnit) AddRpc(mac string, rpc chan interface{}) {
	c.RpcCh[mac] = rpc
}

func (c *CommUnit) SendRpcResponse(mac string, msg interface{}) {
	for k, v := range c.RpcCh {
		if 0 == strings.Compare(k, mac) {
			v <- msg
		}
	}
}

func (c *CommUnit) RetriveConn(mac string) *net.Conn {
	return c.ConnMap[mac]
}

func (c *CommUnit) RetriveMacByConn(conn *net.Conn) string {
	for k, v := range c.ConnMap {
		if v == conn {
			return k
		}
	}
	return ""
}

func CommInit() {
	//	RpcResponse = make(chan interface{}, 1)
	//	ConnMap = make(map[string]net.Conn)
	Comm = NewCommUnit()
}
