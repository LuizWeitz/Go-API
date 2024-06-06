package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/config"
	"github.com/luizweitz/go-api/internal/controllers"
	"github.com/luizweitz/go-api/internal/models"
	"github.com/luizweitz/go-api/internal/repositories"
	"github.com/luizweitz/go-api/internal/services"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.LoadEnvVariables()
	db = config.ConnectDB()

}

func main() {

	db.Table("users").AutoMigrate(&models.User{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/users", userController.Create)
		v1.GET("/users", userController.GetAll)
		v1.GET("/users/:id", userController.GetById)
		v1.PUT("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)
	}

	router.Run(":8080")

}
