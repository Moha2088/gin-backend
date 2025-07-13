package main

import (
	"gin-backend/docs"
	"gin-backend/internal/config"
	"gin-backend/internal/controllers"
	"gin-backend/internal/repositories"
	"gin-backend/internal/routers"
	"gin-backend/internal/services"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config.LoadEnv()
	logger := config.SetupLogger()

	vaultUri := os.Getenv("VaultUri")
	credential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		logger.Warn("Failed to create credential")
	}

	client, err := azsecrets.NewClient(vaultUri, credential, nil)

	if err != nil {
		logger.Warn("Failed to create client")
	}

	dbConfig := config.NewDatabaseConfig(logger, client)

	repository := repositories.NewProjectRepository(logger, dbConfig.GetDatabase())
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
