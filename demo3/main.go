package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index ....")
	name, ok := c.Get("name")
	if !ok {
		name = ""
	}
	time.Sleep(time.Millisecond * 10)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "index",
		"name": name,
	})
}

//定义中间件 统计耗时
func m1(c *gin.Context) {
	log.Println("m1 in ....")
	//计时
	//调用后续的处理函数
	start := time.Now()
	c.Next()
	//阻止调用后续函数
	//c.Abort()
	cost := time.Since(start)
	log.Printf("cost:%v\n", cost)
	log.Println("m1 out ....")
}
func m2(c *gin.Context) {
	log.Println("m2 in ....")
	//计时
	//调用后续的处理函数
	//c.Next()
	c.Set("name", "zhuyijun")
	//阻止调用后续函数
	//c.Abort()
	log.Println("m2 out ....")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库 。。。。
	return func(c *gin.Context) {
		if doCheck {
			log.Println("开始判断是否登录")
			//可以用于鉴权
			_, ok := c.Get("name")
			if !ok {
				log.Println("没有登录")
				c.Abort()
			}
			log.Println("已经登录成功")
			c.Next()
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(m1, m2)
	r.GET("/", authMiddleware(false), indexHandler)
	//第一种
	shopGroup := r.Group("/shop", authMiddleware(true))
	{
		shopGroup.GET("/index", indexHandler)
	}
	//第二种
	vedioGroup := r.Group("/vedio")
	vedioGroup.Use(authMiddleware(true))
	{
		vedioGroup.GET("/index", indexHandler)
	}
	r.Run(":9029")
}
