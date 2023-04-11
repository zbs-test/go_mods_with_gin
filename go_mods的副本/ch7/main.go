package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// user := r.Group("/user")
	// user.GET("/index", func(c *gin.Context) {})
	// user.POST("/login", func(c *gin.Context) {})
	// r.Run()

	r := gin.Default()
	// g1 /user
	user := r.Group("/user")
	// /user/index
	user.GET("/index", func(c *gin.Context) {})
	user.POST("/login", func(c *gin.Context) {})
	// g2 /user/pwds/pwd1
	pwd:=user.Group("/pwds")
	pwd.GET("/pwd1",func(c *gin.Context) {})
	r.Run()

}