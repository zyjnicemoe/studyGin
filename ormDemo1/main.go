package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string `gorm:"default:'男'"`
	Hobby  string
}

func (u UserInfo) TableName() string {
	return "user_info"
}
func init() {

}
func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//创建表
	db.AutoMigrate(&UserInfo{})
	user := &UserInfo{Id: 1}
	log.Println(db.NewRecord(&user))
	////创建数据行
	//user = &UserInfo{Id:2,Name: "zhuyijun",Hobby: "游戏"}
	//db.Create(user)
	db.First(user)
	log.Println(*user)
	////更新
	db.Model(user).Update("hobby", "动漫1")
	log.Println(*user)

	//db.Delete(&user)

}
