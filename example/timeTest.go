package main

import (
	"fmt"
	"time"
)

const base_format = "2006-01-02 15:04:05"

//Parse("2016-12-04 15:39:06 +0800 CST")

func main() {
	//	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Clock())
	//	fmt.Println(time.Now().Unix())
	//	fmt.Println(time.Now().Format(base_format))

	//转换为时间格式字符串
	t := time.Now()
	str_time := t.Format(time.RFC3339)
	fmt.Printf("now time string:%v\n", str_time)
	fmt.Println("Now0:", t.Unix())

	//时间字符串转换为日期格式
	parse_str_time, _ := time.Parse(time.RFC3339, str_time)
	fmt.Printf("string to datetime :%v\n", parse_str_time)
	fmt.Println("Now1:", parse_str_time.Unix())
	fmt.Println("Now1:", parse_str_time.Format(time.RFC3339))
	//fmt.Println("After:", parse_str_time.Unix()-t.Unix())

	local, _ := time.LoadLocation("Asia/Chongqiong")
	//local, _ := time.LoadLocation("Local")
	p, _ := time.ParseInLocation(time.RFC3339, str_time, local)
	fmt.Println("Now2:", p.Format(time.RFC3339))
	fmt.Println("Now2:", p.Unix())

	//fmt.Printf("Year, Week: %v", p.ISOWeek())
	Year, week := p.ISOWeek()
	fmt.Println("Year:", Year, "week:", week)
	fmt.Println("YearDay", p.YearDay())
	zone, offset := p.Zone()
	fmt.Println("Zone:", zone, "Offset:", offset)

	t0 := time.Now()
	duration := t0.Sub(t)
	fmt.Println("Duration %v", duration, time.Since(t))

	complex, _ := time.ParseDuration("10m4s")
	hours, _ := time.ParseDuration("10h")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Println(complex.Seconds())

	old := time.Unix(1534504315, 0)

	fmt.Println("Until:", time.Until(old).Minutes())

	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v \n", now)
	}
	//从unix时间戳转换为字符串
	//	t1 := time.Unix(1534400468, 0)
	//	fmt.Println("Time from Unix():", t1.Local().Format(base_format))
}
