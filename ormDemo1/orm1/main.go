package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//定义模型
type User struct {
	//内嵌gorm.Model
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	//创建索引 addr
	Address  string `gorm:"index:addr"`
	IgnoreMe int    `gorm:"-"`
}

//唯一指定表明
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_user"
	}
	return "user"
}

type Animal struct {
	Id   int64 `gorm:"primary_key"`
	Name string
	Age  int64
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sms_" + defaultTableName
	}
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//禁用复数
	db.SingularTable(true)
	//创建表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	//使用User结构体创建名叫admin的表
	db.Table("admin").CreateTable(&User{})

}
