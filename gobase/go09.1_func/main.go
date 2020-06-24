package main

import "fmt"

//函数 是组织好的、可重复使用的、用于执行指定任务的代码块
//定义一个不需要参数也没有返回值的函数sayHello
func sayHello() {
	fmt.Println("hello big man")
}

//定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("hello", name)
}

//定义接受多个参数的函数并且有一个返回值
func intSum(a, b int) int {
	ret := a + b
	return ret
}

//返回值命名
func intSum2(a, b int) (ret int) {
	ret = a + b
	return
}

//函数接受可变参数，在参数名后面加上...表示可变参数
//可变参数在函数体中是切片类型
func intSum3(a ...int) int {
	fmt.Println(a)        //[20 30]
	fmt.Printf("%T\n", a) //[]int
	ret := 0
	for _, arg := range a {
		ret += arg
	}
	return ret
}
//固定参数和可变参数同时出现时，可变参数要放在最后
func intSum4(a int , b ...int) int {
	ret := a
	for _, arg := range b {
		ret += arg
	}
	return ret
}
func main() {
	//函数调用
	sayHello()
	name := "小公主"
	sayHello2(name)
	sayHello2("小王子")
	r := intSum(10, 20)
	fmt.Println(r)
	r2 := intSum2(30, 40)
	fmt.Println(r2)
	intSum3(20, 30)
	fmt.Println("--------------")
	ret := intSum3()
	ret1 := intSum3(10)
	ret2 := intSum3(10, 20)
	ret3 := intSum3(10, 20, 30)
	fmt.Println(ret)
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(ret3)
	fmt.Println("无情的分割线-------------------")
	ret4 := intSum4(10) // a=10 b=[]
	ret5 := intSum4(10, 20)// a=10 b=[20]
	ret6 := intSum4(10, 20, 30)// a=10 b=[20,30]
	fmt.Println(ret4)
	fmt.Println(ret5)
	fmt.Println(ret6)
}
