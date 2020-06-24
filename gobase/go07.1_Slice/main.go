package main

import "fmt"

//切片 Slice
func main() {
	var a []string
	var b []int

	var c = []bool{true, false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//1.基于数组得到切片
	d := [5]int{55, 56, 57, 58, 59}
	e := d[1:4]
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	//2.切片再次切片
	f := e[:] //e[0:len(e)]
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	//3.make函数构造切片
	g := make([]int, 5, 10) //长度是5,容量是10
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	//通过len()函数获取切片的长度
	fmt.Println(len(g))
	//通过cap()函数获取切片的容量
	fmt.Println(cap(g))

}
