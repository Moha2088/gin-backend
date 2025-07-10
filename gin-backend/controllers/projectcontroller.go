package controllers

import (
	"gin-backend/gin-backend/cqrs/commands"
	"gin-backend/gin-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) *projectController {
	return &projectController{service: service}
}

// @Summary: Creates a new projects
// @Success 201 {ProjectDto}

func (s *projectController) CreateProject(c *gin.Context) {
	var command commands.CreateProjectCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := s.service.CreateProject(command)
	c.IndentedJSON(http.StatusCreated, dto)
}
