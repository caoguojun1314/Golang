package main

import (
	"fmt"
	"reflect"
)

// reflect value

func reflectValue(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n",v,v)
	k := v.Kind() //拿到值对应的类型种类
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		//把反射取到的值转换成一个Float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v,%T\n",ret,ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n",ret,ret)
	}
}
func reflectSetValue(x interface{}){
	v := reflect.ValueOf(x)
	//Elem()用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}
func main() {
	var a int32 = 100
	reflectValue(a)
	var b float32 = 1.234
	reflectValue(b)
	//set value
	var c int32 = 10
	reflectSetValue(&c)
	fmt.Println(c)

}
