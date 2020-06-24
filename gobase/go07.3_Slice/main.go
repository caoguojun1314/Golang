package main

import (
	"fmt"
	"sort"
)

//切片Slice
//切片要初始化之后才能使用
func main() {
	//切片的扩容
	var a []int //此时它并没有申请内存
	//a[0]=100
	//fmt.Println(a) 报错 index out of range [0] with length 0
	for i := 0; i < 10; i++ {
		a = append(a, i) //append自动扩容，可能会改变底层的数组，需要一个变量接收返回值
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a)
	}
	//可以在append中一次追加多个元素
	var b []int
	b = append(b, 1, 2, 3, 4, 5)
	fmt.Println(b)
	//追加一个切片
	c := []int{12, 13, 14, 15}
	b = append(b, c...) //...代表c中的元素一个个拆出来全部追加进去
	fmt.Println(b)
	//切片的copy
	d := []int{1, 2, 3, 4, 5}
	e := make([]int, 5, 5)
	f := e //指向同一个底层数组
	copy(e, d)
	e[0] = 100
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	//切片删除元素
	g := []string{"北京", "上海", "广州", "深圳"}
	g = append(g[0:2], g[3:]...) //别忘记...
	fmt.Println(g)
	//删除索引为index的元素
	//append(g[:index], g[index+1:]...)
	//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	var h = [...]int{3, 7, 8, 9, 1}
	sort.Ints(h[:]) //h[:]得到的是一个切片，指向了底层数组h
	fmt.Println(h)
}
