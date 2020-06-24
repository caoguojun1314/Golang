package main

import "fmt"

//变量

func main() {
	//标准申明格式
	var name string
	var age int
	var isOk bool
	fmt.Print(name, age, isOk)
	//批量申明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Print(a, b, c, d)
	//变量申明同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Print(name1, age1)
	var name2, age2 = "Guojun", 18
	fmt.Print(name2, age2)
	//类型推倒
	var name3 = "小公举"
	var age3 = 17
	fmt.Print(name3, age3)
	//短变量申明
	m := 10
	fmt.Print(m)
	//匿名变量
	/*
	1.函数外的每个语句必须以关键字开头(var, func, const 等)
	2.:= 不能使用在函数外
	3. 匿名变量_多用于站位，表示忽略值
	 */

}
