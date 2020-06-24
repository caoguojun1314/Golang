package main

import (
	"fmt"
	"strings"
)

//map练习
func main() {
	//统计一个字符串中每个单词出现的次数
	//"how do you do"中每个单词出现的次数
	//0.定义一个map[string]int
	var s = "how do you do"
	var wordsCount = make(map[string]int, 10)
	//1.字符串中都有哪些单词
	words := strings.Split(s, " ")
	//2.遍历单词做统计
	for _, word := range words {
		v, ok := wordsCount[word]
		if ok {
			//map中有这个单词的统计
			wordsCount[word] = v + 1
		} else {
			//map中没有这个单词的统计
			wordsCount[word] = 1
		}
	}
	for k, v := range wordsCount {
		fmt.Println(k, v)
	}

}
