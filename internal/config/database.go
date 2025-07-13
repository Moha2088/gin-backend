package config

import (
	"context"
	"gin-backend/internal/models"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	logger       *zap.Logger
	secretClient *azsecrets.Client
}

func NewDatabaseConfig(logger *zap.Logger, secretClient *azsecrets.Client) *DatabaseConfig {
	return &DatabaseConfig{logger: logger, secretClient: secretClient}
}

func (c *DatabaseConfig) GetDatabase() *gorm.DB {

	latestSecretVersion := ""
	secretResponse, err := c.secretClient.GetSecret(context.TODO(), "DBConnection", latestSecretVersion, nil)

	if err != nil {
		c.logger.Warn("Failed to get connectionstring from Key Vault")
	}

	connectionString := *secretResponse.Value
	c.logger.Info(connectionString)
	db, err := gorm.Open(postgres.Open(connectionString))

	if err != nil {
		c.logger.Warn("Error connecting to database!")
	}

	err = db.AutoMigrate(&models.Project{})

	if err != nil {
		c.logger.Warn("Migration failed!")
	}

	return db
}
