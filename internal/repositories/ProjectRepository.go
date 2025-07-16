package repositories

import (
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gin-backend/internal/models"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/errors"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

type ProjectRepository interface {
	CreateProject(command commands.CreateProjectCommand) (dtos.ProjectDto, error)
	GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error)
	GetProjects(query queries.GetAllProjectsQuery) ([]dtos.ProjectDto, error)
	UpdateProject(id uint, command commands.UpdateProjectCommand) (dtos.ProjectDto, error)
	DeleteProject(command commands.DeleteProjectCommand) error
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

func (r *projectRepository) CreateProject(command commands.CreateProjectCommand) (dtos.ProjectDto, error) {

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

	if createdEntity.RowsAffected == 0 {
		return dtos.ProjectDto{}, errors.New("Error creating project!")
	}
	r.logger.Sugar().Info("Rows affected: %d", createdEntity.RowsAffected)

	return project.ToDto(), nil
}

func (r *projectRepository) GetProject(query queries.GetProjectQuery) (dtos.ProjectDto, error) {

	var project models.Project
	response := r.db.Find(&project, query.ProjectId)

	if response.RowsAffected == 0 {
		r.logger.Warn("Error getting project by id")
		return dtos.ProjectDto{}, errors.New("Project not found!")
	}

	return project.ToDto(), nil
}

func (r *projectRepository) GetProjects(query queries.GetAllProjectsQuery) ([]dtos.ProjectDto, error) {
	var projects []models.Project

	response := r.db.Find(&projects)

	if response.RowsAffected == 0 {
		return make([]dtos.ProjectDto, 0), errors.New("Error getting projects")
	}

	var dtos []dtos.ProjectDto

	for _, project := range projects {
		dtos = append(dtos, project.ToDto())
	}

	return dtos, nil
}

func (r *projectRepository) UpdateProject(id uint, command commands.UpdateProjectCommand) (dtos.ProjectDto, error) {
	var project models.Project

	response := r.db.Find(&project, id)

	if response.RowsAffected == 0 {
		return dtos.ProjectDto{}, errors.New("Project not found!")
	}

	project.Name = command.Name
	project.Description = command.Description
	project.Participants = command.Participants
	project.From = command.From
	project.To = command.To

	updateResponse := r.db.Model(&project).Updates(project)

	if updateResponse.RowsAffected == 0 {
		return dtos.ProjectDto{}, errors.New("Error updating project")
	}

	return project.ToDto(), nil
}

func (r *projectRepository) DeleteProject(command commands.DeleteProjectCommand) error {
	var project models.Project
	response := r.db.Find(&project, command.ProjectId)

	if response.RowsAffected == 0 {
		return errors.New("Project not found!")
	}

	deleteResponse := r.db.Delete(&project)

	if deleteResponse.RowsAffected == 0 {
		return errors.New("Error deleting project")
	}
	return nil
}
