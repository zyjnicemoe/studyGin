package routers

import (
	"github.com/gin-gonic/gin"
	"studyGin/bubble/controller"
)

func V1(e *gin.Engine) {
	//v1
	v1Group := e.Group("/v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.AddTodo)
		//查看所有
		v1Group.GET("/todo", controller.FindAll)
		//通过id查看
		v1Group.GET("/todo/:id", controller.FindById)
		//修改
		v1Group.PUT("/todo/:id", controller.Update)
		//删除
		v1Group.DELETE("/todo/:id", controller.Delete)
	}
}
