package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"../proto"
)

const (
	MaxRead = 4096
)

type Login struct {
	Cmd   string `json:"cmd"`
	SeqId string `json:"seqId"`
	Mac   string `json:"mac"`
	Token string `json:"token"`
}

func LoginFunc() ([]byte, uint32) {

	login := Login{
		Cmd:   "login",
		SeqId: "xxxxx",
		Mac:   "00782fe82e35", //Mac:   "aabbccddeeff",
		Token: "xxxx",
	}

	msg, err := json.Marshal(login)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg, proto.CmdKV["login"]
}

func HeartBeat() ([]byte, uint32) {
	heart := proto.ReqParam{
		Cmd:   "heartbeat",
		SeqId: "xxxxx",
		Mac:   "00782fe82e35",
	}
	msg, err := json.Marshal(heart)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg, proto.CmdKV["heartbeat"]
}

//send Rcl request to server
func Rcl() ([]byte, uint32) {
	rcc := proto.ReqParam{
		Cmd:   "rcl",
		SeqId: "xxxxx",
		Mac:   "00782fe82e35",
	}
	msg, err := json.Marshal(rcc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg, proto.CmdKV["rcl"]
}

//send ConfigRead response to server
func ConfigRead() ([]byte, uint32) {
	msg, err := ioutil.ReadFile("configread.json")
	if err != nil {
		fmt.Println("config read failed", err.Error())
	}
	return msg, proto.CmdKV["config_read_resp"]
}

func Verification() ([]byte, uint32) {
	ver := proto.VerificationReq{
		Cmd:         "verification_req",
		SeqId:       "xxxxx",
		Mac:         "00782fe82e35",
		TerminalMac: "myiPhoneMac",
	}
	msg, err := json.Marshal(ver)
	if err != nil {
		Log(err.Error())
	}

	Log(string(msg))
	return msg, proto.CmdKV["verification_req"]
}

func SendData(conn net.Conn, msg []byte, Id uint32) {
	conn.Write(proto.PacketLemon3((msg), Id))
}

func ProtoCycle(conn net.Conn) {
	msg, id := ConfigRead()

	SendData(conn, msg, id)
}
func sendMessage() {
	conn, err := net.Dial("tcp", "192.168.0.12:37001")

	if err != nil {
		panic("Error")
	}
	count := 0
	//Send goroutine, send request
	go sender(conn)
	for {
		//Send message to tcp server
		switch count {
		case 0:
			msg, id := LoginFunc()
			SendData(conn, msg, id)
			count = 1
			break
		case 1:
			msg, id := Rcl()
			SendData(conn, msg, id)
			count = 2
			break
		case 2:
			msg, id := Verification()
			SendData(conn, msg, id)
			count = 3
			break

		}

		fmt.Println("next read data...")

		rbuf := make([]byte, MaxRead+1)
		length, err := conn.Read(rbuf[0 : MaxRead+1])

		if err != nil {
			fmt.Println("Fuck reading ", err.Error())
			time.Sleep(time.Second * 5)
			return
		}
		rbuf[MaxRead] = 0

		readerChannel := make(chan []byte, 4096)
		go reader(readerChannel, conn)
		tmpBuffer := make([]byte, 0)
		tmpBuffer = proto.UnpackLemon3(append(tmpBuffer, rbuf[:length]...), readerChannel)

		time.Sleep(time.Second * 2)

	}
}
func reader(readerChannel chan []byte, conn net.Conn) {
	select {
	case data := <-readerChannel:
		msg := make(map[string]interface{})
		Log("Receive message: ", string(data))
		err := json.Unmarshal(data, &msg)

		//fmt.Println("err:", err, msg)
		if err == nil && msg["cmd"] != nil {
			if msg["cmd"].(string) == "reboot_req" {
				resp := "{\"cmd\":\"reboot_resp\",\"seqId\":\"1234321\",\"code\":\"000\",\"Message\":\"message\"}"
				SendData(conn, []byte(resp), proto.CmdKV["reboot_resp"])
				Log("Send response to server")
			} else if msg["cmd"].(string) == "config_read_req" {
				msg, id := ConfigRead()
				SendData(conn, msg, id)
				Log("send config read response to server")
			} else if msg["cmd"].(string) == "notification_req" {
				resp := proto.RespParam{
					Cmd:   "notification_resp",
					Code:  "111",
					SeqId: "unique id",
				}
				msg, err := json.Marshal(resp)
				if err != nil {
					Log(err.Error())
				}
				SendData(conn, msg, proto.CmdKV["notification_resp"])
			} else if msg["cmd"].(string) == "dns_bogus_write_req" {
				resp := "{\"cmd\":\"dns_bogus_write_resp\", \"seqId\":\"1234321\",\"code\":\"000\",\"Message\":\"message\"}"
				SendData(conn, []byte(resp), proto.CmdKV["dns_bogus_write_resp"])
				Log("Send response to server")
			}
		}

		if strings.Contains(string(data), "heartbeat") {
			go sender(conn)
		}
	}
}

func sender(conn net.Conn) {
	select {
	case <-time.After(time.Second * 30): //30s 超时退出select
		Log("heartbeat...")
		msg, id := HeartBeat()
		SendData(conn, msg, id)
	}
	Log("######################### end of sender ############")
}
func Log(v ...interface{}) {
	fmt.Println(v...)
}

func main() {
	fmt.Println("TCP Client Entity")

	proto.InitKeyValue()

	sendMessage()
	select {}
}
