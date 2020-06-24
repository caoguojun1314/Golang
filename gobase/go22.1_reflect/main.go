package main

import (
	"fmt"
	"reflect"
)

//reflect demo
func reflectType(x interface{}){
	//不知道别人调用这个函数时传入什么类型的变量
	//1.通过类型断言
	//2.借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) //*reflect.rtype
}
type Cat struct{}
type Dog struct{}
func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	//结构体类型
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)
	//Slice
	//Go语言的反射中像数组，切片，Map，指针等类型的变量，它们的.Name()都是返回空
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

}
