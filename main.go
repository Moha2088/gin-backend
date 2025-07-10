package main

import (
	"gin-backend/gin-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	logger := config.SetupLogger()

}
