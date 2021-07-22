package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"studyGin/bubble/dao"
	"studyGin/bubble/models"
	"studyGin/bubble/routers"
)

func init() {
	//链接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
}
func main() {
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})
	//路由
	r := routers.Router()
	r.Run(":9029")
}
