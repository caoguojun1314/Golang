package main

import (
	"fmt"
	"time"
)

//Time
func main() {
	now := time.Now() //时间对象
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	//时间戳：从1970.1.1到现在的一个秒数
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1, timestamp2)
	//将时间戳转换成具体的时间格式
	t := time.Unix(now.Unix(), 0 )
	fmt.Println(t)
	//时间间隔
	n := 5 // type int
	time.Sleep(time.Duration(n)*time.Second)
	fmt.Println("over")
}
