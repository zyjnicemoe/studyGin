package models

import (
	"log"
	"studyGin/bubble/dao"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (t Todo) TableName() string {
	return "todo"
}

func CreateTodo(todo *Todo) error {
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

func FindAll() (*[]Todo, error) {
	todoList := &[]Todo{}
	if err := dao.DB.Find(todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}

func FindById(id string) (*Todo, error) {
	todo := &Todo{}
	if err := dao.DB.Where("id=?", id).Find(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
