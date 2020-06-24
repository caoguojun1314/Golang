package main

import "fmt"

//if判断
func main() {
	//1.基本写法
	score := 65
	if score > 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//2.if判断的特殊写法
	if score2 := 65; score2 > 90{ //score2只在if语句中可以用
		fmt.Println("A")
	} else if score2 > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//3.if语句可以没有else if也可以没有else语句

}
