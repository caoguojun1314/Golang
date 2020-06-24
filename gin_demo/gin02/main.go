package main

import(
	"github.com/gin-gonic/gin"
)
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"hello golang!",
	})
}
func main() {
	r := gin.Default() //返回一个默认路由引擎
	//指定用户使用GET请求访问hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)
	//启动服务
	r.Run(":9090")
}

