package main

import (
	"github.com/sparky0520/go-crud/initializers"
	"github.com/sparky0520/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
