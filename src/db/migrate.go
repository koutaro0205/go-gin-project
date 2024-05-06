package main

import (
	"go-gin-project/initializers"
	model "go-gin-project/models"
)

func main() {
	initializers.ConnectToDB().AutoMigrate(&model.Todo{})
}
