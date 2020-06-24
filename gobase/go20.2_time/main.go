package main

import (
	"fmt"
	"time"
)

// Time demo
func main() {
	now := time.Now()
	fmt.Println(now)
	//now + 1 Hour
	t := now.Add(time.Hour)
	fmt.Println(t)
	//Sub
	fmt.Println(t.Sub(now))
	//定时器
	/*for tmp := range time.Tick(time.Second*2){
		fmt.Println(tmp)
	}
	*/
	//时间格式化 y    m  d  h  m  s
	//         2006 01 02 15 04 05
	ret1 := now.Format("2006-01-02")
	fmt.Println(ret1)
	ret2 := now.Format("2006.01.02 15:04:05")
	fmt.Println(ret2)
	ret3 := now.Format("2006.01.02 03:04:05 PM")
	fmt.Println(ret3)
	ret4 := now.Format("2006.01.02 03:04:05.000 PM")
	fmt.Println(ret4)
	//解析字符串类型的时间
	timeStr := "2019/08/07 15:00:00"
	//1.拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.根据时区去解析一个字符串格式的时间
	timeObj1, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("Parse timeStr failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj1)
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("Parse timeStr failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj2)
}
