// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// 	"github.com/gin-gonic/gin"
// )
// // 钩子函数
// //定义一个中间键m1统计请求处理函数耗时,查看控制台(后台)的请求耗时
// func m1(c *gin.Context) {
// 	fmt.Println("m1 in...")
// 	start := time.Now()
// 	// c.Next() //调用后续的处理函数
// 	c.Abort()//阻止调用后续的处理函数
// 	cost := time.Since(start)
// 	fmt.Printf("cost:%v\n", cost)
// }

// func index(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"msg": "ok",
// 	})
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/", m1, index)
// 	r.Run()
// }

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		tmp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "tmp"，这一点很重要
			log.Println("Done! in path " + tmp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	r.Run()
}