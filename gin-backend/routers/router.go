package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1/projects")
	{
		v1.POST("/")
	}

	return router
}
