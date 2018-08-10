package server

import (
	"encoding/json"
	"time"

	log "../log"
	"../proto"
)

/// Reboot/Notification/config_read etc
//RPC Command
func ExecSendCommand(msg string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(msg), &dat); err == nil {
		cmd := dat["cmd"].(string)
		mac := dat["mac"].(string)
		//Find cmd from CMDKV binary value
		cmdId := proto.CmdKV[cmd]

		//Find net.Conn of mac
		conn := Comm.RetriveConn(mac)
		rpc := make(chan interface{})
		Comm.AddRpc(mac, rpc)
		defer close(rpc)

		log.Debug("RPC Command:[%s]", msg)
		//Send cmd to mac
		_, err := (*conn).Write(proto.PacketLemon3([]byte(msg), cmdId))

		if err != nil {
			return "error in send command", err
		}

		select {
		case m := <-rpc:
			log.Debug("exec send command, %s", m.(string))
			return m.(string), nil
		case <-time.After(5 * time.Second):
			bmsg, err := json.Marshal([]byte("\"cmd\":\"failed\""))
			log.Debug("exec send command timeout, %s", string(bmsg))
			return string(bmsg), err
		}
	} else {
		return "invalid json", err
	}
}
