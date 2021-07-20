package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//name := c.Query("query")
		//name := c.DefaultQuery("query","somebody")
		name, ok := c.GetQuery("query")
		age := c.Query("age")
		if !ok {
			//取不到
			name = "somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9029")
}
