package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/config"
	docs "github.com/luizweitz/go-api/docs"
	"github.com/luizweitz/go-api/internal/controllers"
	"github.com/luizweitz/go-api/internal/models"
	"github.com/luizweitz/go-api/internal/repositories"
	"github.com/luizweitz/go-api/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.LoadEnvVariables()
	db = config.ConnectDB()

}

// @contact.name Luiz Weitz
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	docs.SwaggerInfo.Title = "Go API Master Example"
	docs.SwaggerInfo.Description = "This a project is a example of API Rest in Go using Gin, Gorm & PostgreSQL."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	db.Table("users").AutoMigrate(&models.User{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/users", userController.Create)
		v1.PUT("/users", userController.Update)
		v1.GET("/users", userController.GetAll)
		v1.GET("/users/:id", userController.GetById)
		v1.DELETE("/users/:id", userController.Delete)
	}

	urlSwaggerJson := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwaggerJson))

	router.Run(":8080")

}
