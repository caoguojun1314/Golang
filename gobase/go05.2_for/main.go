package main

import "fmt"

//for循环
func main() {
	//1.for循环的基本写法
	for i:= 0; i<10; i++ {
		fmt.Println(i)
	}
	//2.省略初始语句，但是必须保留初始语句后面的分号
	j := 0
	for ; j <10 ; j++{
		fmt.Println(j)
	}
	//3.省略初始语句，结束语句
	var k = 10
	for k > 0 {
		fmt.Printf("k:%v\n", k)
		k--
	}
	//4.无线循环
	//for {
	//	fmt.Println("hello 小王子")
	//}
	//5.break跳出for循环
	for l:=0; l<10; l++{
		fmt.Println(l)
		if l==3 {
			break
		}
	}
	//6.continue继续下一次循环
	for i:=0; i<10; i++{
		if i==3 {
			continue // 跳过本次for循环，继续下一次for循环
		}
		fmt.Println(i)

	}
}
