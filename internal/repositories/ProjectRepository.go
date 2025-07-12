package repositories

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gin-backend/internal/models"
	"time"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type ProjectRepository interface {
	CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto
	GetProject(query queries.GetProjectQuery) dtos.ProjectDto
	GetProjects(query queries.GetAllProjectsQuery) []dtos.ProjectDto
	UpdateProject(command commands.UpdateProjectCommand) dtos.ProjectDto
	DeleteProject(command commands.DeleteProjectCommand)
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

func (r *projectRepository) CreateProject(command commands.CreateProjectCommand) dtos.ProjectDto {

	var isActive bool
	var isCompleted bool
	currentDate := time.Now()

	if currentDate.After(command.From) && currentDate.Before(command.To) {
		isActive = true
	}

	if currentDate.After(command.To) {
		isCompleted = true
	}

	project := models.Project{
		Model:        gorm.Model{},
		Name:         command.Name,
		Description:  command.Description,
		Participants: command.Participants,
		IsActive:     isActive,
		From:         command.From,
		To:           command.To,
		IsCompleted:  isCompleted,
		CreatedAt:    currentDate,
	}

	createdEntity := r.db.Create(&project)
	r.logger.Sugar().Info("Rows affected: %d", createdEntity.RowsAffected)

	return project.ToDto()
}

func (r *projectRepository) GetProject(query queries.GetProjectQuery) dtos.ProjectDto {
	panic("unimplemented")
}

func (r *projectRepository) GetProjects(query queries.GetAllProjectsQuery) []dtos.ProjectDto {
	panic("")
}

func (r *projectRepository) UpdateProject(command commands.UpdateProjectCommand) dtos.ProjectDto {
	panic("")
}

func (r *projectRepository) DeleteProject(command commands.DeleteProjectCommand) {
	panic("")
}
