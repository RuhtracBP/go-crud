package main

import (
	"github.com/RuhtracBP/go-crud/initializers"
	"github.com/RuhtracBP/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
