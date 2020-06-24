package main

import "fmt"

//map(映射)
func main() {
	//只声明map类型，但是没有初始化，a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil)
	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	//map中添加键值对
	a["小王子"] = 18
	a["小公举"] = 16
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)
	//声明map，同时完成初始化
	b := map[int]bool{
		1: true,
		0: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("b:%T\n", b)

	//var c map[int]int
	//c[100] = 200 //c这个map没有初始化不能直接操作,panic: assignment to entry in nil map
	//fmt.Println(c)

	//判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["小王子"] = 99
	scoreMap["小公主"] = 98
	//判断恶魔在不在scoreMap中
	value, ok := scoreMap["恶魔"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("恶魔在scoreMap中", value)
	} else {
		fmt.Println("查无此人")
	}
	value, ok = scoreMap["小王子"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("小王子在scoreMap中", value)
	} else {
		fmt.Println("查无此人")
	}
	//遍历map
	// for range 遍历
	//map是无序的，键值对和添加的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//只遍历map中的key
	for k := range scoreMap {
		fmt.Println(k)
	}
	//只遍历map中的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}
	//delete 删除键值对
	delete(scoreMap, "小公主")
	fmt.Println(scoreMap)
}
