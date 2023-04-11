package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./views/login.html", "./views/index.html") //加载页面
	// login.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	})
	// index.html
	r.POST("/", func(c *gin.Context) {
		username := c.PostForm("username") //对应h5表单中的name字段
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	
	r.Run()
}