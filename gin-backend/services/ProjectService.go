package services

import (
	"gin-backend/gin-backend/cqrs/commands"
	"gin-backend/gin-backend/cqrs/queries"
	"gin-backend/gin-backend/dtos"
	"gin-backend/gin-backend/repositories"
)

type ProjectService interface {
	CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto
	GetProject(query queries.GetProjectQuery) dtos.ProjectDto
	GetProjects()
	UpdateProject(command commands.UpdateProjectCommand) dtos.ProjectDto
	DeleteProject(command commands.DeleteProjectCommand)
}

type projectService struct {
	repository repositories.ProjectRepository
}

func NewProjectService(repository repositories.ProjectRepository) ProjectService {
	return &projectService{
		repository: repository,
	}
}

func (p *projectService) CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto {
	project := p.repository.CreateProject(command)
	return project
}

// DeleteProject implements ProjectService.
func (p *projectService) DeleteProject(command commands.DeleteProjectCommand) {
	p.repository.DeleteProject(command)
}

// GetProject implements ProjectService.
func (p *projectService) GetProject(query queries.GetProjectQuery) dtos.ProjectDto {
	project := p.repository.GetProject(query)
	return project
}

// GetProjects implements ProjectService.
func (p *projectService) GetProjects() {
	panic("unimplemented")
}

// UpdateProject implements ProjectService.
func (p *projectService) UpdateProject(command commands.UpdateProjectCommand) dtos.ProjectDto {
	project := p.repository.UpdateProject(command)
	return project
}
