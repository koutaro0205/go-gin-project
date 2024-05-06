package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbInit() *gorm.DB {
	// dbを作成
	// NOTE: "@tcp(XX)"の"XX"には、docker-compose.ymlのサービス名orコンテナ名が入る
	dsn := "root:password@tcp(db)/go_gin_project_db?charset=utf8mb4&parseTime=true&loc=Local"
	// DBインスタンスを初期化
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

type Todo struct {
	gorm.Model

	Title string `json: "title"`
	Done  bool   `json: "done"`
	Body  string `json: "body"`
}

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrate
	db.AutoMigrate(&Todo{})

	router := gin.Default()

	// 全件取得
	router.GET("/todos", func(c *gin.Context) {
		todos := []Todo{}
		result := db.Find(&todos)

		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"todos":  todos,
		})
	})

	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// データを登録
		result := db.Create(&todo)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 正常系レスポンス
		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todo": todo})
	})

	// 更新
	router.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		todo := Todo{} // フロントから送信されたデータがバインドされる

		prevTodo := db.First(&todo, "id = ?", id) // 先に更新するデータを取得
		if errors.Is(prevTodo.Error, gorm.ErrRecordNotFound) {
			log.Fatal(prevTodo.Error)
			return
		}

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Save(&todo)
		if err := result.Error; err != nil {
			log.Fatal(err)
			return
		}

		// 正常系レスポンス
		c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": todo})
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		result := db.Where("id = ?", id).Delete(&Todo{})
		if err := result.Error; err != nil {
			log.Fatal(err)
			return
		}

		// 正常系レスポンス
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully", "id": id})
	})

	router.Run(":8000")
}
