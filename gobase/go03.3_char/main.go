package main

import "fmt"

//字符
func main() {
	//byte uint8的别名 ASCII码
	//rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n",c1, c2)

	s:="hello沙河"
	fmt.Println(len(s))
	for i:=0; i<len(s); i++{ //byte
		fmt.Printf("%c\n",s[i])
	}
	fmt.Println(s)
	for _, r := range s { //rune
		fmt.Printf("%c\n",r)
	}
}
