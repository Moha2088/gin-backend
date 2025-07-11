package config

import (
	"gin-backend/internal/models"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	logger *zap.Logger
}

func NewDatabaseConfig(logger *zap.Logger) *DatabaseConfig {
	return &DatabaseConfig{logger: logger}
}

func (c *DatabaseConfig) GetDatabase() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		c.logger.Info("Error loading env file!")
	}

	connectionString := os.Getenv("DBConnectionString")

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
