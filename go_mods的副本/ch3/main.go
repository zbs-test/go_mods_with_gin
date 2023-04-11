package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	// get index and return index.html
	r.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"name":"admin",
			"pwd":"123456",
		})
	})
	
	r.GET("/test",func(c *gin.Context){
		c.HTML(http.StatusOK,"test.html",gin.H{
			"name":"null",
		})
	})
	r.Run()
}