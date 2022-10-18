package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建默认路由
	router := gin.Default()
	//2.绑定路由规则,执行函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello！欢迎来到GO世界！")
	})
	//绑定路由规则 不加函数,表示空白页
	router.POST("/xxxPost")
	router.GET("/xxxPost")
	router.PUT("/xxx")
	// 3.默认端口是8080,也可以指定端口 r.Run(":80")
	router.Run()
}
