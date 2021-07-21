package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"strconv"
	"studyGin/bubble/dao"
	"studyGin/bubble/models"
)

func CreateTodo(todo *models.Todo) error {
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(todo).Error; err != nil {
		log.Fatalln("创建清单失败")
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func main() {
	//链接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	//告诉滚、框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架 去哪里找模板文件
	r.LoadHTMLGlob("templates/**")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			todo := &models.Todo{}
			err := c.BindJSON(todo)
			if err != nil {
				log.Fatalln("获取todo数据失败")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			err = CreateTodo(todo)
			if err != nil {
				log.Fatalln("创建代办失败")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, todo)
		})
		//查看所有
		v1Group.GET("/todo", func(c *gin.Context) {
			todoList := &[]models.Todo{}
			if err := dao.DB.Find(todoList).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error,
				})
				return
			}
			log.Println(*todoList)
			c.JSON(http.StatusOK, *todoList)
		})
		//通过id查看
		v1Group.GET("/todo/:id", func(c *gin.Context) {
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
		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusInternalServerError, "无效id")
				return
			}
			todo := &models.Todo{}
			if err := dao.DB.Where("id=?", id).Find(todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.BindJSON(todo)
			if err := dao.DB.Save(todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, todo)
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			ID, _ := strconv.Atoi(id)
			if !ok {
				c.JSON(http.StatusInternalServerError, "无效id")
				return
			}
			dao.DB.Delete(&models.Todo{Id: ID})
			c.JSON(http.StatusOK, "删除成功")
		})
	}

	r.Run(":9029")
}
