package server

import (
	"fmt"
	"net"

	log "../log"
	"../proto"
)

// Handle Each client request
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	log.Info("Connection from: %s", connFrom)

	m := func(conn net.Conn) {
		err := conn.Close()
		log.Info("Closing connection: %s\n", connFrom)
		checkError(err, "Close:")
	}
	defer m(conn)

	//talktoclients(conn, "{\"result\":\"ok\"}")

	for {
		ibuf := make([]byte, MaxRead+1)

		length, err := conn.Read(ibuf[0 : MaxRead+1])
		//checkError(err, "Read From client failed")
		if err != nil {
			log.Error("ERROR %s\n", err.Error())
			return
		}
		ibuf[MaxRead] = 0
		//		fmt.Println("length=", length)

		readerChannel := make(chan []byte, 16)

		//reader get the parsed message from
		go reader(readerChannel, conn)

		//TODO: TCP粘包问题需要解决
		tmpBuffer := make([]byte, 0)
		tmpBuffer = append(tmpBuffer, ibuf[:length]...)
		for {
			tmpBuffer = proto.UnpackLemon3(tmpBuffer, readerChannel)
			if len(tmpBuffer) > 0 {
				tmpBuffer = proto.UnpackLemon3(tmpBuffer, readerChannel)
			} else {
				break
			}
		}
	}
}

func reader(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			//Process data to get Msg and cmdId
			msg, cmdId := handleMsg(data, conn)

			if cmdId == proto.CmdKV["reboot_resp"] ||
				cmdId == proto.CmdKV["upgrade_resp"] ||
				cmdId == proto.CmdKV["cc_write_resp"] ||
				cmdId == proto.CmdKV["rc_write_resp"] ||
				cmdId == proto.CmdKV["upgrade_resp"] ||
				cmdId == proto.CmdKV["config_read_resp"] ||
				cmdId == proto.CmdKV["notification_resp"] {
				//send message to rpc command.
				fmt.Println("Msg:", msg)
				//RpcResponse <- msg
				Comm.SendRpcResponse(Comm.RetriveMacByConn(&conn), msg)
			}

			if cmdId == 0 {
				continue
			}
			//Send response to client
			talktoclients(conn, msg, cmdId)
		}
	}
}

func talktoclients(to net.Conn, msg string, cmdId uint32) {
	//data := "Hello Client"
	wrote, err := to.Write(proto.PacketLemon3([]byte(msg), cmdId))
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
	if cmdId == proto.CmdKV["heartbeat"] {
		log.Info("This is heartbeat")
	}
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error())
	}
}
