package main

import "fmt"

//函数进阶之变量作用域
//定义全局变量num
var num = 10

//定义函数
func testGlobal() {
	num := 100
	name := "Big toy"
	//可以在函数中使用变量
	//1.先在自己函数中找，找到了就用自己函数中的变量值
	//2.函数中找不到就往外层找全局变量
	fmt.Println("变量num", num)
	fmt.Println(name)
}
func main() {
	testGlobal()
	//外层不能访问函数的内部变量（局部变量）
	//fmt.Println(name)

	//变量i只在for循环的语句块中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i) 外层访问不到内部for语句块中的变量

	//函数可以作为变量
	fmt.Println("无情的吃货线-------------")
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()

}
