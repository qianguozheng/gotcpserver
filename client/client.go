package main

import (
	"fmt"
	"net"
	"time"

	"encoding/json"

	"../proto"
)

const (
	MaxRead = 1024
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
		Mac:   "aabbccddeeff",
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
		Mac:   "aabbccddeeff",
	}
	msg, err := json.Marshal(heart)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg, proto.CmdKV["heartbeat"]
}

func Rcl() ([]byte, uint32) {
	rcc := proto.ReqParam{
		Cmd:   "rcl",
		SeqId: "xxxxx",
		Mac:   "aabbccddeeff",
	}
	msg, err := json.Marshal(rcc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg, proto.CmdKV["rcl"]
}

func SendData(conn net.Conn, msg []byte, Id uint32) {
	conn.Write(proto.PacketLemon3((msg), Id))
}

func ProtoCycle(conn net.Conn) {

	//msg, id := HeartBeat()

	//msg, id := LoginFunc()
	msg, id := Rcl()
	SendData(conn, msg, id)
}
func sendMessage() {
	conn, err := net.Dial("tcp", "192.168.0.12:37001")

	if err != nil {
		panic("Error")
	}

	for {
		//		words := "{\"cmd\":\"login\",\"seqId\":\"1234321\",\"Message\":\"message\"}"
		//		conn.Write(proto.PacketLemon3([]byte(words), 0x34))

		//Send message to tcp server
		ProtoCycle(conn)

		fmt.Println("Send Data Already")

		rbuf := make([]byte, MaxRead+1)
		length, err := conn.Read(rbuf[0 : MaxRead+1])
		if err != nil {
			fmt.Println("Fuck reading ", err.Error)
			return
		}
		rbuf[MaxRead] = 0

		readerChannel := make(chan []byte, 16)
		go reader(readerChannel)
		tmpBuffer := make([]byte, 0)
		tmpBuffer = proto.UnpackLemon3(append(tmpBuffer, rbuf[:length]...), readerChannel)

		time.Sleep(time.Second * 2)

	}
}
func reader(readerChannel chan []byte) {
	select {
	case data := <-readerChannel:
		fmt.Println("Message: ", string(data))
	}
}

func main() {
	fmt.Println("TCP Client Entity")

	currency := 1
	count := 1

	proto.InitKeyValue()

	for i := 0; i < currency; i++ {
		go func() {
			for j := 0; j < count; j++ {
				sendMessage()
			}
		}()
	}
	select {}
}
