package server

import (
	"encoding/json"
	"fmt"
	"time"

	"../proto"
)

/// Reboot, etc
//RPC Command
func ExecSendCommand(msg string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(msg), &dat); err == nil {
		cmd := dat["cmd"].(string)
		mac := dat["mac"].(string)
		//Find cmd from CMDKV binary value
		cmdId := proto.CmdKV[cmd]

		//Find net.Conn of mac
		conn := ConnMap[mac]

		//Send cmd to mac
		_, err := conn.Write(proto.PacketLemon3([]byte(msg), cmdId))

		if err != nil {
			return "error in send command", err
		}

		select {
		case m := <-RpcResponse:
			fmt.Println("exec send command", m.(string))
			return m.(string), nil
		case <-time.After(5 * time.Second):
			bmsg, err := json.Marshal([]byte("\"cmd\":\"failed\""))
			fmt.Println("exec send command", bmsg)
			return string(bmsg), err
		}
	} else {
		return "invalid json", err
	}
}
