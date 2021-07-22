package routers

import (
	"github.com/gin-gonic/gin"
	"studyGin/bubble/controller"
)

func V2(e *gin.Engine) {
	//v1
	v1Group := e.Group("/v2")
	{
		//查看所有
		v1Group.GET("/ok", controller.Get)
	}
}
