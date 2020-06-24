package main

import (
	"fmt"
)

//数组相关内容
//数组的长度一定常量
func main() {
	var a [3]int //长度是数组类型的一部分
	a[0] = 3
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	//数组的初始化
	//1.定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	fmt.Println(cityArray[3])
	//2.编译器推倒数组长度
	var boolArray = [...]bool{true, false, true, false, false}
	fmt.Println(boolArray)
	//3.使用索引值的方式初始化
	var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	fmt.Println(langArray)
	fmt.Printf("%T\n", langArray)
	//4.数组的遍历
	var cityArray2 = [...]string{"包头", "呼市", "南宁", "天津"}
	//4.1 for 循环遍历数组
	for i := 0; i < len(cityArray2); i++ {
		fmt.Println(cityArray2[i])
	}
	//4.2 for range 遍历
	fmt.Println("for range 遍历")
	for index, value := range cityArray2 {
		fmt.Println(index, value)
	}
	for _, value := range cityArray {
		fmt.Println(value)
	}
	//二维数组
	//多维数组只有第一层可以使用[...]来让编译器推导数组的长度
	cityArray3 := [4][2]string{
		{"北京", "上海"},
		{"重庆", "成都"},
		{"东邪", "西毒"},
		{"南帝", "北丐"},
	}
	fmt.Println(cityArray3)
	fmt.Println(cityArray3[0][1])
	//二维数组的遍历
	fmt.Println("二维数组的遍历")
	for _, v1 := range cityArray3 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//数组是值类型
	e := [3]int{1, 2, 3}
	fmt.Println(e)
	modify(e)
	fmt.Println(e)
	f := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(f)
	modify2(f)
	fmt.Println(f)
}
func modify(a [3]int) {
	a[0] = 100
}
func modify2(a [3][2]int) {
	a[0][0] = 100
}
