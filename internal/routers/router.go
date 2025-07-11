package routers

import (
	"gin-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c controllers.ProjectController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1/projects")
	{
		v1.POST("/", c.CreateProject)
		v1.GET("/:id", c.GetProject)
		v1.GET("/", c.GetProjects)
		v1.PUT("/:id", c.UpdateProject)
		v1.DELETE("/:id", c.DeleteProject)
	}

	return router
}
