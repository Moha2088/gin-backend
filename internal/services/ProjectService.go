package services

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gin-backend/internal/repositories"
)

type ProjectService interface {
	CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto
	GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error)
	GetProjects(query queries.GetAllProjectsQuery) []dtos.ProjectDto
	UpdateProject(id uint, command commands.UpdateProjectCommand) dtos.ProjectDto
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

func (p *projectService) DeleteProject(command commands.DeleteProjectCommand) {
	p.repository.DeleteProject(command)
}

func (p *projectService) GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error) {
	project, err := p.repository.GetProject(query)

	if err != nil {
		return project, err
	}
	return project, nil
}

func (p *projectService) GetProjects(query queries.GetAllProjectsQuery) []dtos.ProjectDto {
	panic("unimplemented")
}

func (p *projectService) UpdateProject(id uint, command commands.UpdateProjectCommand) dtos.ProjectDto {
	project := p.repository.UpdateProject(id, command)
	return project
}
