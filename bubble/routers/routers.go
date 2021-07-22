package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	//告诉滚、框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架 去哪里找模板文件
	r.LoadHTMLGlob("templates/**")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	V1(r)
	V2(r)
	return r
}
