package service

import (
	"go-gin-project/graph/model"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]*model.User, error) {
	users := []*model.User{}
	db.Find(&users)

	return users, nil
}

func GetUser(db *gorm.DB, id string) (*model.User, error) {
	user := model.User{}

	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, input model.CreateUserInput) (*model.User, error) {
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	db.Create(&user)
	return &user, nil
}
