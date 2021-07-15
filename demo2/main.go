package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {

	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取请求参数
		//name := c.Query("query")
		//name := c.DefaultQuery("query","参数获取不到啦")

		name,ok := c.GetQuery("query")
		age := c.Query("age")
		if !ok {
			name = "somebody"
		}
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	r.Run(":8080")
}
