package repositories

import (
	"gin-backend/gin-backend/cqrs/commands"
	"gin-backend/gin-backend/cqrs/queries"
	"gin-backend/gin-backend/dtos"

	"go.uber.org/zap"
)

type ProjectRepository interface {
	CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto
	GetProject(query queries.GetProjectQuery) dtos.ProjectDto
	GetProjects()
	UpdateProject(commands.UpdateProjectCommand) dtos.ProjectDto
	DeleteProject(commands.DeleteProjectCommand)
}

type projectRepository struct {
	logger *zap.Logger
}

func NewProjectRepository(logger *zap.Logger) ProjectRepository {
	return &projectRepository{
		logger: logger,
	}
}

// DeleteProject implements ProjectRepository.
func (r *projectRepository) DeleteProject(commands.DeleteProjectCommand) {
	panic("unimplemented")
}

// GetProject implements ProjectRepository.
func (r *projectRepository) GetProject(query queries.GetProjectQuery) dtos.ProjectDto {
	panic("unimplemented")
}

// GetProjects implements ProjectRepository.
func (r *projectRepository) GetProjects() {
	panic("unimplemented")
}

// UpdateProject implements ProjectRepository.
func (r *projectRepository) UpdateProject(commands.UpdateProjectCommand) dtos.ProjectDto {
	panic("unimplemented")
}

func (r *projectRepository) CreateProject(commands.CreateProjectCommand) dtos.ProjectDto {

}
