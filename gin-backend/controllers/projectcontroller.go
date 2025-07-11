package controllers

import (
	"gin-backend/gin-backend/cqrs/commands"
	"gin-backend/gin-backend/cqrs/queries"
	"gin-backend/gin-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController interface {
	CreateProject(c *gin.Context)
	GetProject(c *gin.Context)
	GetProjects(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type projectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) ProjectController {
	return &projectController{service: service}
}

// CreateProject godoc
// @Summary: Creates a new projects
// @Success 201 {object} dtos.ProjectDto
// @Failure 400
// @Tags projects
// @Router /api/v1/projects/ [post]
func (s *projectController) CreateProject(c *gin.Context) {
	var command commands.CreateProjectCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := s.service.CreateProject(command)
	c.IndentedJSON(http.StatusCreated, dto)
}

// GetProject godoc
// @Summary: Gets a project by ID
// @Success 200 {object} dtos.ProjectDto
// @Failure 400
// @Tags projects
// @Router /api/v1/projects/{id}/ [get]
func (s *projectController) GetProject(c *gin.Context) {

	if id := c.Param("id"); id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
	}

	uuid_string := c.Param("id")
	uuid, err := uuid.Parse(uuid_string)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse project ID"})
	}

	getProjectQuery := queries.GetProjectQuery{ProjectId: uuid}
	project := s.service.GetProject(getProjectQuery)

	c.IndentedJSON(http.StatusOK, project)
}

func (s *projectController) GetProjects(c *gin.Context) {
	panic("unimplemented")
}

func (s *projectController) UpdateProject(c *gin.Context) {

}

// DeleteProject godoc
// @Summary: Deletes a project by ID
// @Success 204
// @Failure 400
// @Tags projects
// @Router /api/v1/projects/{id}/ [delete]
func (s *projectController) DeleteProject(c *gin.Context) {
	var command commands.DeleteProjectCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.service.DeleteProject(command)
	c.Status(http.StatusNoContent)
}
