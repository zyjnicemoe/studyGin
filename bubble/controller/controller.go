package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"studyGin/bubble/models"
)

func AddTodo(c *gin.Context) {
	todo := &models.Todo{}
	err := c.BindJSON(todo)
	if err != nil {
		log.Fatalln("获取todo数据失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.CreateTodo(todo)
	if err != nil {
		log.Fatalln("创建代办失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func FindById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "无效id")
		return
	}
	todo := &models.Todo{}
	todo, err := models.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}
