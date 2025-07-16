package services

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gin-backend/internal/repositories"
)

type ProjectService interface {
	CreateProject(command commands.CreateProjectCommand) (dtos.ProjectDto, error)
	GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error)
	GetProjects(query queries.GetAllProjectsQuery) ([]dtos.ProjectDto, error)
	UpdateProject(id uint, command commands.UpdateProjectCommand) (dtos.ProjectDto, error)
	DeleteProject(command commands.DeleteProjectCommand) error
}

type projectService struct {
	repository repositories.ProjectRepository
}

func NewProjectService(repository repositories.ProjectRepository) ProjectService {
	return &projectService{
		repository: repository,
	}
}

func (p *projectService) CreateProject(command commands.CreateProjectCommand) (dtos.ProjectDto, error) {
	project, err := p.repository.CreateProject(command)
	return project, err
}

func (p *projectService) DeleteProject(command commands.DeleteProjectCommand) error {
	err := p.repository.DeleteProject(command)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectService) GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error) {
	project, err := p.repository.GetProject(query)
	return project, err
}

func (p *projectService) GetProjects(query queries.GetAllProjectsQuery) ([]dtos.ProjectDto, error) {
	projects, err := p.repository.GetProjects(query)
	return projects, err
}

func (p *projectService) UpdateProject(id uint, command commands.UpdateProjectCommand) (dtos.ProjectDto, error) {
	project, err := p.repository.UpdateProject(id, command)
	return project, err
}
