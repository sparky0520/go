package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sparky0520/go-crud/controllers"
	"github.com/sparky0520/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/add", controllers.PostsCreate)
	r.GET("/all", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("/update/:id", controllers.PostsUpdate)
	r.DELETE("/delete/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
