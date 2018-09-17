package server

import (
	"net"
	"strings"

	"sync"

	"../log"
)

/*
	Internal Communication Structure
*/

type CommUnit struct {
	ConnMap map[string]net.Conn
	RpcCh   map[string]chan interface{}
	sync.Mutex
}

var Comm *CommUnit

func NewCommUnit() *CommUnit {
	comm := &CommUnit{
		ConnMap: make(map[string]net.Conn),
		RpcCh:   make(map[string]chan interface{}),
	}
	return comm
}

func (c *CommUnit) AddConn(mac string, conn net.Conn) {
	c.Lock()
	defer c.Unlock()
	c.ConnMap[mac] = conn
}

func (c *CommUnit) AddRpc(mac string, rpc chan interface{}) {
	c.Lock()
	defer c.Unlock()
	c.RpcCh[mac] = rpc
}

func (c *CommUnit) SendRpcResponse(mac string, msg interface{}) {
	c.Lock()
	defer c.Unlock()
	for k, v := range c.RpcCh {
		log.Debug("k=%s, mac=%s", k, mac)
		if 0 == strings.Compare(k, mac) {
			v <- msg
		}
	}
}

func (c *CommUnit) RetriveConn(mac string) net.Conn {
	c.Lock()
	defer c.Unlock()
	return c.ConnMap[mac]
}

func (c *CommUnit) RetriveMacByConn(conn net.Conn) string {
	c.Lock()
	defer c.Unlock()
	for k, v := range c.ConnMap {
		//log.Debug("conn=%v, v=%v", *conn, *v)
		if (v) == (conn) {
			//log.Debug("RetriveMacByConn: %s", k)
			return k
		}
	}
	log.Debug("Not found mac by conn")
	return ""
}

func CommInit() {
	//	RpcResponse = make(chan interface{}, 1)
	//	ConnMap = make(map[string]net.Conn)
	Comm = NewCommUnit()
}
