package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询学生信息成功",
		})
	})
	r.POST("/create_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建学生信息成功",
		})
	})
	r.PUT("/update_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新学生信息成功",
		})
	})
	r.DELETE("/delete_student", func(c *gin.Ccurlontext) {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除学生信息成功",
		})
	})
	r.Run()
}