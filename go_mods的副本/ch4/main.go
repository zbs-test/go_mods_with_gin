package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		pwd := c.Query("pwd")
		fmt.Printf("name:%s ; pwd:%s",name,pwd)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
	r.Run()
}