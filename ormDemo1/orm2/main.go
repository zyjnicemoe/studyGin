package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	Id     uint `gorm:"AUTO_INCREMENT"`
	Name   string
	Gender string `gorm:"default:'男'"`
	Hobby  string
}

func (u User) TableName() string {
	return "test_user"
}
func main() {

	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.LogMode(true)
	//db.SetLogger(gorm.Logger{revel.TRACE})
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//user := &User{}
	db.AutoMigrate(&User{})

	//user = &User{Id: 1,Name: "张三",Hobby: "游戏"}
	//log.Println("开启事务")
	//tx := db.Begin()
	//if err := db.Create(user).Error; err != nil {
	//	log.Println("创建失败")
	//	//事务回滚
	//	defer tx.Rollback()
	//	log.Println("回滚事务完成")
	//	return
	//}
	//db.Commit()
	//log.Println("事务提交")

	//var count int
	//err = db.Model(&User{}).Where(&User{Id: 1}).Count(&count).Error
	//if err != nil {
	//	log.Println("查询失败")
	//	return
	//}
	//log.Println(count)

	//user = &User{Name: "张三",Hobby: "游戏"}
	//log.Println("开启事务")
	//tx := db.Begin()
	//if err := db.Create(user).Error; err != nil {
	//	log.Println("创建失败")
	//	//事务回滚
	//	defer tx.Rollback()
	//	log.Println("回滚事务完成")
	//	return
	//}
	//db.Commit()
	//log.Println("事务提交")

	user := &User{}
	db.First(user)
	log.Println(&*user == user)
	log.Println(*user)

}
