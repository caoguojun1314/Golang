package main

import "fmt"

//常量

//const pi = 3.1415
//const e = 2.7182
//const (
//	pi = 3.1415
//	e= 2.7182
//)
const (
	n1 = 10
	n2
	n3
)
//iota在const关键字出现时将被重置为0，const中每新增一行常量申明将使iota计数一次
const(
	a1 = iota
	a2
	a3
	a4
)
//使用匿名变量_跳过某些值
const(
	b1 = iota
	b2
	_
	b3
)
//iota申明中间插队
const(
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)
const c5 = iota

func main() {
	fmt.Print(n1, n2, n3)
	fmt.Print(a1, a2, a3, a4)
	fmt.Print(b1, b2, b3)
	fmt.Print(c1,c2,c3,c4,c5)
}
