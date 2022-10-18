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
		//可通过contxt.Query获取带参数的路由
		param := c.Query("userName")
		//可通过contxt.DefaultQuery如果无参数则给默认值
		param1 := c.DefaultQuery("userName", "默认值")
		//返回浏览器 状态码,输出
		c.String(http.StatusOK, "获取参数:"+param+"\n")
		c.String(http.StatusOK, "获取参数:"+param1+"\n")
		c.String(http.StatusOK, "Hello！欢迎来到GO世界！\n")
		c.String(http.StatusOK, "请求路径:"+c.Request.URL.Path+"\n")
		c.String(http.StatusOK, "请求ip:"+c.ClientIP())
	})

	//注意: :name是映射匹配    使用*name是模糊匹配
	router.GET("/:name", func(c *gin.Context) {
		//可通过contxt.Query获取带参数的路由
		param := c.Param("name")
		//返回浏览器 状态码,输出
		c.String(http.StatusOK, "获取参数:"+param+"\n")
	})

	router.POST("/testPost", testPost)

	//绑定路由规则 不加函数,表示空白页
	router.GET("/xxxPost")
	router.PUT("/xxx")
	// 3.默认端口是8080,也可以指定端口 r.Run(":80")
	router.Run()
}

/*
*
获取表单参数
*/
func testPost(c *gin.Context) {
	//获取表单参数
	form := c.PostForm("name")
	c.String(http.StatusOK, "获取表单参数:"+form)
}
