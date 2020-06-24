package split

import "strings"

//split.go
//Spilt将s按照sep进行分割，返回一个字符串的切片
//Split("我爱你"，"爱") => ["我", "你"]
//中文占3-4个字节所以idx+1 改为 idx+ len（sep）
func Split(s, sep string)(ret []string){
	ret = make([]string, 0 , strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1{
		ret = append(ret, s[:idx])
		s = s[idx + len(sep) :]
		idx = strings.Index(s, sep)
	}
	ret = append (ret, s)
	return
}
