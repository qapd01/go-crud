package main

import (
	"github/qapd01/go-crud/controllers"
	"github/qapd01/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB() 
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
