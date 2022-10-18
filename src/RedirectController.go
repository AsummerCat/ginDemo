package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
重定向
*/
func redirectTest(c *gin.Context) {
	//状态码 301 重定向地址 支持内部地址和外部地址
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	//c.Redirect(http.StatusMovedPermanently, "/")
}
