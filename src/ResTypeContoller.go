package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 gin的多种返回渲染格式
*/

/*
返回json数据
*/
func resJson(c *gin.Context) {
	badJson := gin.H{"name": "小明", "age": 18}
	//以json的形式返回
	c.JSON(http.StatusOK, badJson)

	//结构体响应
	var msg struct {
		Name   string
		Number int
	}
	c.JSON(http.StatusOK, msg)
}

/*
返回html模板数据
*/
func resHtml(c *gin.Context) {
	//1.首先需要在路由那边加载模板文件
	/*
		router := gin.Default()
		//加载模板文件
		router.LoadHTMLGlob("templates/*")
	*/
	//最终json将title替换
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我的标题"})
}
