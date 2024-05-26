package service

import (
	"go-gin-project/graph/model"

	"gorm.io/gorm"
)

func GetAllTodos(db *gorm.DB) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	db.Preload("User").Find(&todos)

	return todos, nil
}

func GetTodo(db *gorm.DB, id string) (*model.Todo, error) {
	todo := model.Todo{}

	if err := db.Preload("User").First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func CreateTodo(db *gorm.DB, input model.CreateTodoInput) (*model.Todo, error) {
	todo := model.Todo{
		UserID: input.UserID,
		Title:  input.Title,
		Done:   false,
	}
	// エラーハンドリングを行う
	db.Create(&todo)

	return &todo, nil
}

func UpdateTodo(db *gorm.DB, input model.UpdateTodoInput) (*model.Todo, error) {
	todo := model.Todo{}
	if err := db.First(&todo, "id = ?", input.ID).Error; err != nil {
		return nil, err
	}

	todo.Title = input.Title
	todo.Done = input.Done

	db.Save(&todo)

	return &todo, nil
}

func DeleteTodo(db *gorm.DB, id string) (*model.Todo, error) {
	todo := model.Todo{}
	if err := db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}

	db.Delete(&todo)

	return &todo, nil
}
