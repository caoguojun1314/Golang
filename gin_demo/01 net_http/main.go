package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	//_, _ =fmt.Fprint(w, "<h1>Hello Golang!</h1><h2> i wanna kiss you!</h2>")
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ =fmt.Fprint(w, string(b) )
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v \n", err)
	}
}
