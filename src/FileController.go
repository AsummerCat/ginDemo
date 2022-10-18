package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/*
多文件上传
*/
func uploadMultipart(c *gin.Context) {
	fileList := c.Request.MultipartForm

	//读取其中的文件
	files := fileList.File["files"]
	for _, file := range files {
		log.Print(file.Filename)
	}
}

/*
单文件上传
*/
func upload(c *gin.Context) {
	//从表单中提取文件
	file, _ := c.FormFile("file")
	log.Print(file.Filename)
}
