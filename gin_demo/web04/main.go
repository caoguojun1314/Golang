package main

import(
	"fmt"
	"net/http"
	"html/template"
)
//遇事不决，先写注释
func sayHello(w http.ResponseWriter, r *http.Request){
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse file failed, err:%v\n", err)
		return
	}
	//渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v\n", err)
		return
	}

}
