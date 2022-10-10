package main

import (
	"github/qapd01/go-crud/initializers"
	"github/qapd01/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
  