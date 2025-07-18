package main

import (
	"context"
	"gin-backend/docs"
	"gin-backend/internal/config"
	"gin-backend/internal/controllers"
	"gin-backend/internal/repositories"
	"gin-backend/internal/routers"
	"gin-backend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config.LoadEnv()
	logger := config.SetupLogger()

	secretClientconfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientconfig.GetSecretClient())

	client := secretClientconfig.GetSecretClient()
	latestSecretVersion := ""
	response, err := client.GetSecret(context.TODO(), "DBConnection", latestSecretVersion, nil)

	if err != nil {
		logger.Warn("Failed to get secret from Key Vault")
	}

	connectionString := response.Value

	repository := repositories.NewProjectRepository(logger, dbConfig.GetDatabase(*connectionString))
	service := services.NewProjectService(repository)
	controller := controllers.NewProjectController(service)

	router := routers.SetupRouter(controller)

	router.Use(cors.Default())

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
