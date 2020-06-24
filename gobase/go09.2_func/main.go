package main

import "fmt"

//Go语言中参数类型简写
func intSum(a, b int) int {
	ret := a + b
	return ret
}

//定义具有多个返回值的函数,多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main() {
	r := intSum(10, 20)
	fmt.Println(r)
	r2, r3 := calc(200, 100)
	fmt.Println(r2, r3)

}
