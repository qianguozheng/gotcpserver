package main

import (
	"fmt"
	"time"
)

const base_format = "2006-01-02 15:04:05"

func main() {
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format(base_format))

	//转换为时间格式字符串
	t := time.Now().UTC()
	str_time := t.Format(base_format)
	fmt.Printf("now time string:%v\n", str_time)
	fmt.Println("Now:", t.Unix())

	//时间字符串转换为日期格式
	//time.Sleep(time.Second * 2)
	//fmt.Println("Current Now:", time.Now().Unix)
	parse_str_time, err := time.Parse(base_format, str_time)
	fmt.Println("Err:", err)
	fmt.Printf("string to datetime :%v\n", parse_str_time)
	fmt.Println("Now:", parse_str_time.Unix())
	//fmt.Println("After:", parse_str_time.Unix()-t.Unix())
	//	local, _ := time.LoadLocation("Asia/Chongqiong")
	//	p, _ := time.ParseInLocation(base_format, str_time, local)
	//	fmt.Println("Now:", p.Unix())

	t1 := time.Unix(1534400468, 0) //Parse("2016-12-04 15:39:06 +0800 CST")
	fmt.Println("Time from Unix():", t1.Local().Format(base_format))
}
