package main

import (
	"gin-backend/docs"
	"gin-backend/gin-backend/config"
	"gin-backend/gin-backend/controllers"
	"gin-backend/gin-backend/repositories"
	"gin-backend/gin-backend/routers"
	"gin-backend/gin-backend/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	logger := config.SetupLogger()

	repository := repositories.NewProjectRepository(logger)
	service := services.NewProjectService(repository)
	controller := controllers.NewProjectController(service)

	router := routers.SetupRouter(controller)

	// Swagger setup
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Gin Backend API"
	docs.SwaggerInfo.Description = "Backend API in Gin"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8081/swagger/doc.json")),
	)

	// Test Request
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin is running!",
		})
	})

	router.SetTrustedProxies(nil)

	router.Run(":8081")

}
