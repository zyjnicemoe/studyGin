package main

import "github.com/jinzhu/gorm"

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func init() {
	gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
}
func main() {

}
