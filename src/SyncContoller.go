package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/*
同步请求和异步请求
注意 启动goroutime的时候,不应该使用原始上下文,必须使用它的只读副本
*/

func syncTest(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.String(http.StatusOK, "同步请求返回")
}

/*
异步请求
*/
func AsyncTest(c *gin.Context) {
	//创建一个context的只读副本
	context := c.Copy()
	go func() {
		//休眠3秒
		time.Sleep(3 * time.Second)
		log.Print("异步执行:" + context.Request.URL.Path)
	}()
	c.String(http.StatusOK, "异步请求返回")
}
