package main

import (
	"fmt"
	"math"
)

//基本数据类型

func main() {

	//整型
	/*
		//十进制打印成二进制
		n := 10
		fmt.Printf("%d\n", n) //占位符%d表示十进制
		fmt.Printf("%b\n", n) //占位符%b表示二进制
		// 八进制 以0开头
		m := 075
		fmt.Print(m)
		fmt.Printf("%o\n", m)//占位符%o表示八进制
		fmt.Printf("%d\n", m)
		//十六进制 以0x开头
		f := 0xff
		fmt.Printf("%x\n", f)
		fmt.Print(f)
		//uint8 0~255
		var age uint8
		age = 255
		fmt.Print(age)
	*/
	//浮点数
	fmt.Print(math.MaxFloat32)
	fmt.Print(math.MaxFloat64)
	//布尔值
	var a bool
	fmt.Print(a)
	a = true
	fmt.Print(a)
	//布尔类型变量默认的值为false
	//Go语言中不允许将整型强制转换成布尔型
	//布尔型无法参与数值运算，也无法与其他类型进行转换
	//字符串
	s1 := "hello beijing"
	s2 := "你好 北京"
	fmt.Print(s1)
	fmt.Print(s2)
	//转义符
	//打印windows平台下的一个路径c:\code\go.exe
	fmt.Print("c:\\code\\go.exe")
	fmt.Printf("%t制表符\n换行符")
	//多行字符串
	s3 := `
	多行字符串
	两个字符串
	之间的内容
	会原样输出
	\n
	\t
	`
	fmt.Print(s3)

}
