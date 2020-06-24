package main

import "fmt"

//map切片
//引用类型slice，map一定要初始化之后才能使用
func main() {
	//元素类型为map的切片
	var mapSlice = make([]map[string]int, 8) //只完成了切片的初始化
	//[nil nil nil nil nil nil nil nil]
	fmt.Println(mapSlice[0] == nil)
	//还需要完成内部map元素初始化
	mapSlice[0] = make(map[string]int, 8) //完成了map的初始化
	mapSlice[0]["小王子"] = 100
	fmt.Println(mapSlice)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	//值为切片的map
	var sliceMap = make(map[string][]int, 8) //只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		//sliceMap中没有中国这个键
		sliceMap["中国"] = make([]int, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}
	//遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

}
