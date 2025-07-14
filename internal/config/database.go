package config

import (
	"gin-backend/internal/models"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	logger *zap.Logger
}

func NewDatabaseConfig(logger *zap.Logger, secretClient *azsecrets.Client) *DatabaseConfig {
	return &DatabaseConfig{logger: logger}
}

func (c *DatabaseConfig) GetDatabase(connectionString string) *gorm.DB {

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
