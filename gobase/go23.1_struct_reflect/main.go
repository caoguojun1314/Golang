package main

import (
	"fmt"
	"reflect"
)

//struct reflect 结构体反射
type student struct{
	Name string `json:"name" ini:"s_name"`
	Score int  `json:"score" ini:"s_score"`

}
func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 99,
	}
	//通过反射去获取反射体中的所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v type:%v\n", t.Name(),t.Kind())
	//遍历结构体变量的所有字段
	for i:=0;i<t.NumField();i++{
		fileObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v \n", fileObj.Name,fileObj.Type,fileObj.Tag)
		fmt.Println(fileObj.Tag.Get("json"),fileObj.Tag.Get("ini"))
	}
	//根据名字去取结构体中的字段
	fileObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v \n", fileObj2.Name,fileObj2.Type,fileObj2.Tag)
	}
}
