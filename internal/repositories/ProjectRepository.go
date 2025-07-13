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
	UpdateProject(id uint, command commands.UpdateProjectCommand) dtos.ProjectDto
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

	var project models.Project
	response := r.db.Find(&project, query.ProjectId)

	if response.Error != nil {
		r.logger.Warn("Error getting project by id")
	}

	return project.ToDto()
}

func (r *projectRepository) GetProjects(query queries.GetAllProjectsQuery) []dtos.ProjectDto {
	var projects []models.Project

	response := r.db.Find(projects)

	if response.Error != nil {
		r.logger.Warn("Error getting projects")
	}

	var dtos []dtos.ProjectDto

	for _, project := range projects {
		dtos = append(dtos, project.ToDto())
	}

	return dtos
}

func (r *projectRepository) UpdateProject(id uint, command commands.UpdateProjectCommand) dtos.ProjectDto {
	var project models.Project

	response := r.db.Find(&project, id)

	if response.Error != nil {
		r.logger.Warn("Error getting project")
	}

	project.Name = command.Name
	project.Description = command.Description
	project.Participants = command.Participants
	project.From = command.From
	project.To = command.To

	updateResponse := r.db.Model(&project).Updates(project)

	if updateResponse.Error != nil {
		r.logger.Warn("Error updating project")
	}

	return project.ToDto()
}

func (r *projectRepository) DeleteProject(command commands.DeleteProjectCommand) {
	var project models.Project
	response := r.db.Find(&project, command.ProjectId)

	if response.Error != nil {
		r.logger.Warn("Error getting project")
	}

	deleteResponse := r.db.Delete(&project)

	if deleteResponse.Error != nil {
		r.logger.Warn("Error deleting project")
	}
}
