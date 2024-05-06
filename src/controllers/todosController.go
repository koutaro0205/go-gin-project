package controllers

import (
	"errors"
	"go-gin-project/initializers"
	"go-gin-project/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB = initializers.ConnectToDB()

func GetTodoList(c *gin.Context) {
	todos := []models.Todo{}
	result := DB.Find(&todos)

	if err := result.Error; err != nil {
		log.Fatal(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"todos":  todos,
	})
}

func AddTodo(c *gin.Context) {
	todo := models.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// データを登録
	result := DB.Create(&todo)
	if err := result.Error; err != nil {
		log.Fatal(err.Error())
		return
	}

	// 正常系レスポンス
	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todo": todo})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	todo := models.Todo{} // フロントから送信されたデータがバインドされる

	prevTodo := DB.First(&todo, "id = ?", id) // 先に更新するデータを取得
	if err := prevTodo.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal(err.Error())
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := DB.Save(&todo)
	if err := result.Error; err != nil {
		log.Fatal(err)
		return
	}

	// 正常系レスポンス
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	result := DB.Where("id = ?", id).Delete(&models.Todo{})
	if err := result.Error; err != nil {
		log.Fatal(err)
		return
	}

	// 正常系レスポンス
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully", "id": id})
}
