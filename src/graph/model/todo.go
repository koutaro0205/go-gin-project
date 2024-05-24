package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	UserID    string    `json:"userId"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	todo.ID = uuid.New().String()
	return
}
