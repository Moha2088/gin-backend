package controllers

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
		return
	}

	id := c.Param("id")

	getProjectQuery := queries.GetProjectQuery{ProjectId: c.GetUint(id)}
	project, err := s.service.GetProject(getProjectQuery)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, project)
}

// GetProjects godoc
// @Summary: Gets all projects
// @Success 200 {object} []dtos.ProjectDto
// @Failure 404
// @Tags projects
// @Router /api/v1/projects/ [get]
func (s *projectController) GetProjects(c *gin.Context) {
	projects := s.service.GetProjects(queries.GetAllProjectsQuery{})

	if len(projects) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, projects)
}

// UpdateProject godoc
// @Summary: Updates a project by ID
// @Success 200 {object} dtos.ProjectDto
// @Failure 400
// @Tags projects
// @Router /api/v1/projects/{id}/ [put]
func (s *projectController) UpdateProject(c *gin.Context) {
	var command commands.UpdateProjectCommand

	if id := c.Param("id"); id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	id := c.Param("id")

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	project := s.service.UpdateProject(c.GetUint(id), command)
	c.IndentedJSON(http.StatusOK, project)
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
