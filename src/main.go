package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type Todo struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	Done  bool   `json: "done"`
	Body  string `json: "body"`
}

func main() {
	todos := []Todo{{
		Id:    1,
		Body:  "default Body",
		Title: "default Title",
		Done:  false,
	}}

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
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
		// HACK: GORM等を使用してデータベースに登録する
		todo.Id = len(todos) + 1
		todos = append(todos, todo)

		// 正常系レスポンス
		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todos": todos})
	})

	router.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo // フロントから送信されたデータ

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parameterのidに紐づくtodoデータを探す
		_, index, hasTodo := lo.FindIndexOf(todos, func(t Todo) bool {
			return strconv.Itoa(t.Id) == id
		})

		// 該当のデータが存在しない場合、エラーを返す
		if !hasTodo {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "not found"})
			return
		}

		// データを更新
		todos[index].Title = todo.Title
		todos[index].Body = todo.Body
		todos[index].Done = todo.Done

		// 正常系レスポンス
		c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "id": id})
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo // フロントから送信されたデータ

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// // Parameterのidに紐づくtodoデータを探す
		_, index, hasTodo := lo.FindIndexOf(todos, func(t Todo) bool {
			return strconv.Itoa(t.Id) == id
		})

		// 該当のデータが存在しない場合、エラーを返す
		if !hasTodo {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "not found"})
			return
		}

		// データを削除
		todos = append(todos[:index], todos[index+1:]...)

		// 正常系レスポンス
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully", "id": id})
	})

	router.Run(":8000")
}
