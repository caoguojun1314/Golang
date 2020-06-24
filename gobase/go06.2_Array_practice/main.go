package main

import "fmt"

//练习题
func main() {
	//1.求数组[1,3,5,7,8]所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	//把数组的每个元素遍历出来，累加到一个sum变量就可以了
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)
	//2.找出数组中和为指定值的两个元素的下标，
	//比如从数组[1,3,5,7,8]中找出和为8的两个元素元素的下标分别为(0,3)和(1,2)
	//定义两个for循环，外层从第一个开始遍历，内层for循环从外层加1开始遍历
	//它们两个数的和为8
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
