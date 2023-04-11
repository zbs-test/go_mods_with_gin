# 这是go 工程化实践（测试版）
# 使用lapce了
## ch1
- what have I done
1. 使用gin框架,go-1.20
```bash
mkdir ch1 && cd ch1
go mod init ch1
go get -u github.com/gin-gonic/gin
```
2. main.go
```go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//定义路由的GET方法及响应处理函数
	r.GET("/hello", func(c *gin.Context) {
		//将发送的信息封装成JSON发送给浏览器
		c.JSON(http.StatusOK, gin.H{
			//这是我们定义的数据
			"message": "快速入门",
		})
	})
	r.Run() //默认在本地8080端口启动服务
}
```
3. download and run
```bash
go mod tidy
go mod vendor
go run main.go
# another shell
curl localhost:8080/hello
```
# ch2
- what have I done
1. restful api
```bash
GET用来获取资源
POST用来新建资源
PUT用来更新资源
DELETE用来删除资源

 GET	 /student	 查询学生信息
 POST	 /create_student	 创建学生信息
 PUT	 /updata_student	 更新学生信息
 DELETE	 /delete_student	 删除学生信息
```
2. 简单实现restful api(main.go)
3. 使用curl来模拟postman
```bash
# get
curl localhost:8080/student
# post 这里没有要传输的数据, 使用 -F "name=<something>" -F "email=qa@qq.com"来传值
curl localhost:8080/create_student
# put
curl -X PUT localhost:8080/update_student
# delete -v 是显示更多请求的信息
curl -v -X DELETE localhost:8080/delete_student
```
## ch3
- what have I done
1. 前面步骤类似,创建要渲染html页面
```bash
mkdir templates && touch index.html
```
2. 运行和查看/index
3. 自己创建页面和展示,使用数据绑定
## ch4
- what have I done 
1. 获取url请求的参数,参数是使用?和&隔开,这里不再用curl,用浏览器
2. localhost:8080/?name=admin&pwd=123456
3. 页面呈现你的请求参数
## ch5
- what have I done
1. 获得Form表单的请求参数数据
2. 在login.html实现登陆逻辑代码,涉及到input submit,和页面跳转(未实现)
3. 在index.html实现渲染username,password(未实现)
## ch6
- what I have done
1. 实现url动态传参
2. 浏览器输入: localhost:8080/user/<somevalues>
## ch7
- what I have done
1. 路由组的实现,统一处理路由,保持美观
2. curl请求
```bash
curl -X GET localhost:8080/index
curl -X POST localhost:8080/login
```
3. 注释和实现路由嵌套,按标注的方法查询
## ch8
- what I have done
> 中间件：在处理请求的过程中，加入钩子函数，这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登陆认证，权限校验，记录日志等。具体使用方法如下
1. 实现中间件, 做请求耗时统计
2. 注释看第二段代码
> 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。(我不会)
## ch9
- what I have done
1. gin 参数绑定
> 为了能够更方便的获取请求相关参数，提高开发效率，我们可以使用ShouldBind，它能够基于请求自动提取JSON，Form表单，Query等类型的值，并把值绑定到指定的结构体对象，具体使用方法如下
2. 创建查询: localhost:8080/user?username=good&password=123
3. 查看后台的获取数据
## ch10(已略过)
- what I have done
1. 实现简单的文件上传
## ch11
- what I have done
1. 路由重定向和http重定向(就是请求到其他网络页面去)
2. 查看main.go
## ch12(已略过)
- what I have done
1. 设置和获取cookie
## ch13
- what I have done
1. 日志记录到文件
2. 这是信息不再输出到控制台,而是输入了gin.log文件,查看它
## ch14(已失败)
- what I have done
1. 实现热重栽(未实现)
```bash
go get -v -u github.com/pilu/fresh
```
2. 使用命令fresh 代替 go run main.go
## 