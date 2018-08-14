package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"

	log "github.com/qianguozheng/gotcpserver/log"
)

const (
//ConstHeader         = "www.hiweeds.com"
//ConstCommandLength   = 4
//ConstSaveDataLength = 4
)

type RespParam struct {
	Cmd   string   `json:"cmd"`
	Code  string   `json:"code"`
	SeqId string   `json:"seqId"`
	Data  []string `json:"data"`
}

type VerificationData struct {
	TerminalMac string `json:"terminalMac"`
	Valid       int    `json:"valid"`
	AuthType    int    `json:"authType"`
	AuthId      string `json:"authId"`
	UpRate      int    `json:"upRate"`
	DownRate    int    `json:"downRate"`
	TcpLimit    int    `json:"tcpLimit"`
	UdpLimit    int    `json:"udpLimit"`
}

type Verification struct {
	Cmd   string           `json:"cmd"`
	SeqId string           `json:"seqId"`
	Code  string           `json:"code"`
	Data  VerificationData `json:"data"`
}

type VerificationReq struct {
	Cmd         string `json:"cmd"`
	SeqId       string `json:"seqId"`
	Mac         string `json:"mac"`
	TerminalMac string `json:terminalMac`
}

type Notification struct {
	Cmd         string `json:"cmd"`
	SeqId       string `json:"seqId"`
	Mac         string `json:"mac"`
	TerminalMac string `json:"terminalMac"`
	Valid       int    `json:"valid"`
	AuthType    int    `json:"authType"`
	AuthId      string `json:"authId"`
	UpRate      int    `json:"upRate"`
	DownRate    int    `json:"downRate"`
	TcpLimit    int    `json:"tcpLimit"`
	UdpLimit    int    `json:"udpLimit"`
}

type AutoPortalStop struct {
	IOSEnable     int    `json:iOSEnable`
	AndroidEnable int    `json:androidEnable`
	Ios           ReqHdr `json:ios`
	Android       ReqHdr `json:android`
}
type ReqHdr struct {
	Host      []string `json:host`
	Uri       []string `json:uri`
	UserAgent []string `json:userAgent`
}

type ReqParam struct {
	Cmd   string `json:"cmd"`
	SeqId string `json:"seqId"`
	Mac   string `json:"mac"`
}

type WebReadData struct {
	Ver string `json:"ver"`
	Url string `json:"url"`
	Md5 string `json:"md5"`
}
type WebRead struct {
	Cmd   string      `json:"cmd"`
	Code  string      `json:"code"`
	SeqId string      `json:"seqId"`
	Data  WebReadData `json:"data"`
}

type WebWrite struct {
	Cmd   string `json:"cmd"`
	SeqId string `json:"seqId"`
	Mac   string `json:"mac"`
	Ver   int    `json:"ver"`
	Url   string `json:"url"`
	Md5   string `json:"md5"`
}

type Upgrade struct {
	Cmd   string `json:"cmd"`
	SeqId string `json:"seqId"`
	Mac   string `json:"mac"`
	Ver   int    `json:"ver"`
	Url   string `json:"url"`
	Md5   string `json:"md5"`
}

type DnsBogus struct {
	Domain string `json:"domain"`
	Host   string `json:"host"`
}
type DnsBogusWrite struct {
	Cmd   string     `json:"cmd"`
	SeqId string     `json:"seqId"`
	Mac   string     `json:"mac"`
	Bogus []DnsBogus `json:"bogus"`
}

var CmdKV map[string]uint32

func InitKeyValue() {
	CmdKV = make(map[string]uint32)
	//Basic API
	CmdKV["login"] = 0x80000010
	CmdKV["heartbeat"] = 0x80000012
	CmdKV["rcl"] = 0x80000011

	//Config API
	CmdKV["status"] = 0X80000022
	CmdKV["config"] = 0x80000103
	CmdKV["web_read_resp"] = 0X80000106
	CmdKV["web_write_resp"] = 0X80000107
	CmdKV["resource_read_resp"] = 0X80000108
	CmdKV["resource_write_resp"] = 0X80000109
	CmdKV["web_read_req"] = 0X00000106
	CmdKV["web_write_req"] = 0X00000107
	CmdKV["resource_read_req"] = 0X00000108
	CmdKV["resource_write_req"] = 0X00000109
	CmdKV["dns_bogus_write_req"] = 0x00000111
	CmdKV["dns_bogus_write_resp"] = 0x80000111

	//Core API
	CmdKV["verification_req"] = 0x00000013
	CmdKV["verification_resp"] = 0x80000013

	//trigger by server
	CmdKV["notification_req"] = 0x00000101
	CmdKV["notification_resp"] = 0X80000101

	CmdKV["config_read_req"] = 0x00000103
	CmdKV["config_read_resp"] = 0x80000103
	CmdKV["reboot_req"] = 0x00000100
	CmdKV["reboot_resp"] = 0X80000100

	CmdKV["upgrade_req"] = 0x00000102
	CmdKV["upgrade_resp"] = 0X80000102
	//cloud config
	CmdKV["cc_write_req"] = 0x00000104
	CmdKV["cc_write_resp"] = 0X80000104
	//vpn
	CmdKV["rc_write_req"] = 0x00000110
	CmdKV["rc_write_resp"] = 0X80000110
}

func PacketLemon3(messgage []byte, cmd uint32) []byte {
	return append(append(IntToBytes(len(messgage)), UintToBytes(cmd)...), messgage...)
}

func UnpackLemon3(buffer []byte, readerChannel chan []byte) []byte {
	totalLength := buffer[0:4]
	cmdId := buffer[4:8]
	length, err := strconv.ParseInt(fmt.Sprintf("%x", totalLength), 16, 0)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	log.Debug("[FromClient]= [%d] [0x%x] [%s]\n", length, cmdId, string(buffer[8:length+8]))

	readerChannel <- buffer[8 : length+8]

	return buffer[length+8:]
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func UintToBytes(n uint32) []byte {
	x := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToUint(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
