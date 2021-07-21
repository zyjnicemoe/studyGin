package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	url := "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local"
	//链接数据库
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return DB.DB().Ping()
}

func Close() {
	defer DB.Close()
}
