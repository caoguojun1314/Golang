package main

import "fmt"

//defer:延迟执行，延迟处理的语句按defer定义的逆序进行执行
func main() {
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
}
