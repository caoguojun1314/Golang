package main

import "fmt"

//nil
func main() {
	var a []int     //申明int类型切片
	var b = []int{} //申明并且初始化
	var c = make([]int, 0)
	if a == nil {
		fmt.Println("a == nil")
	}
	fmt.Println(a, len(a), cap(a))
	if b == nil {
		fmt.Println("b == nil")
	} else {
		fmt.Println("b !== nil")
	}
	fmt.Println(b, len(b), cap(b))
	if c == nil {
		fmt.Println("c == nil")
	} else {
		fmt.Println("c !== nil")
	}
	fmt.Println(c, len(c), cap(c))
	// 要判断一个切片为空，用len(s) == 0 来判断，不应该使用s == nil 来判断
	//切片的赋值拷贝
	d := make([]int, 3) //[0,0,0]
	e := d              //共用一个底层数组
	e[0] = 100
	fmt.Println(d)
	fmt.Println(e)
	//切片的遍历
	f := []int{1, 2, 3, 4, 5}
	//根据索引来遍历
	for i := 0; i < len(f); i++ {
		fmt.Println(i, f[i])
	}
	//for range 遍历
	fmt.Println("for range 遍历")
	for index, value := range f {
		fmt.Println(index, value)
	}

}
