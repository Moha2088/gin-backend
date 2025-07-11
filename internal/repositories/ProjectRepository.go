package repositories

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gorm.io/gorm"

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
	db     *gorm.DB
}

func NewProjectRepository(logger *zap.Logger, db *gorm.DB) ProjectRepository {
	return &projectRepository{
		logger: logger,
		db:     db,
	}
}

func (r *projectRepository) DeleteProject(commands.DeleteProjectCommand) {
	panic("unimplemented")
}

func (r *projectRepository) GetProject(query queries.GetProjectQuery) dtos.ProjectDto {
	panic("unimplemented")
}

func (r *projectRepository) GetProjects() {
	panic("unimplemented")
}

func (r *projectRepository) UpdateProject(commands.UpdateProjectCommand) dtos.ProjectDto {
	panic("unimplemented")
}

func (r *projectRepository) CreateProject(commands.CreateProjectCommand) dtos.ProjectDto {
	panic("unimplemented")
}
