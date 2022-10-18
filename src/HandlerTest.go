package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handlerAll() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("中间件开始执行了")
		//执行的业务 设置变量到context中
		c.Set("request", "中间件")
		//执行目标函数
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
	}
}

func handlerPortion() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("部分中间件开始执行了")
		//执行的业务 设置变量到context中
		c.Set("request", "中间件")
		//执行目标函数
		c.Next()
		status := c.Writer.Status()
		fmt.Println("部分中间件执行完毕", status)
	}
}
