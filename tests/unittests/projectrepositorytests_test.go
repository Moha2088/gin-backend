package unittests

import (
	"context"
	"gin-backend/internal/config"
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/cqrs/queries"
	"gin-backend/internal/dtos"
	"gin-backend/internal/repositories"
	"gin-backend/tests/fixtures"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateProject_ShouldReturnCreatedProjectdto_WhenCreatingProject(t *testing.T) {
	ctx := context.Background()
	logger := config.SetupLogger()
	secretClientConfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientConfig.GetSecretClient())
	postgresContainer := fixtures.SetupTestContainer(ctx, t)
	connectionString, err := postgresContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	testDb := dbConfig.GetDatabase(connectionString)
	testRepo := repositories.NewProjectRepository(logger, testDb)

	createProjectCommand := commands.CreateProjectCommand{
		Name:         "TestCreateProject",
		Description:  "Testdescription",
		Participants: "Participant",
		From:         time.Now(),
		To:           time.Now().Add(time.Hour),
	}

	response, err := testRepo.CreateProject(createProjectCommand)

	assert.NotEmpty(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, createProjectCommand.Name, response.Name)
	assert.EqualValues(t, createProjectCommand.Description, response.Description)
	assert.EqualValues(t, createProjectCommand.Participants, response.Participants)
	assert.EqualValues(t, createProjectCommand.From, response.From)
	assert.EqualValues(t, createProjectCommand.To, response.To)
	assert.IsType(t, dtos.ProjectDto{}, response)
	assert.NotNil(t, response.ProjectId)

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatal("Cleanup failed: ", err)
		}
	})
}

func Test_GetProject_ShouldReturnProject_WhenProjectExists(t *testing.T) {
	// Arrange
	ctx := context.Background()
	logger := config.SetupLogger()
	secretClientConfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientConfig.GetSecretClient())
	postgresContainer := fixtures.SetupTestContainer(ctx, t)
	connectionString, err := postgresContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	testDb := dbConfig.GetDatabase(connectionString)
	testRepo := repositories.NewProjectRepository(logger, testDb)

	createProjectCommand := commands.CreateProjectCommand{
		Name:         "TestCreateProject",
		Description:  "Testdescription",
		Participants: "Participant",
		From:         time.Now(),
		To:           time.Now().Add(time.Hour),
	}

	// Act
	createResponse, err := testRepo.CreateProject(createProjectCommand)

	// Assert
	assert.NotEmpty(t, createResponse)
	assert.Nil(t, err)
	assert.IsType(t, dtos.ProjectDto{}, createResponse)
	assert.NotNil(t, createResponse.ProjectId)

	getProjectQuery := queries.GetProjectQuery{ProjectId: createResponse.ProjectId}

	getResponse, err := testRepo.GetProject(getProjectQuery)
	assert.Nil(t, err)
	assert.NotNil(t, getResponse)
	assert.IsType(t, dtos.ProjectDto{}, getResponse)

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatal("Cleanup failed: ", err.Error())
		}
	})
}

func Test_GetProject_ShouldReturnError_WhenProjectDoesNotExists(t *testing.T) {
	// Arrange
	ctx := context.Background()
	logger := config.SetupLogger()
	secretClientConfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientConfig.GetSecretClient())
	postgresContainer := fixtures.SetupTestContainer(ctx, t)
	connectionString, err := postgresContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	testDb := dbConfig.GetDatabase(connectionString)
	testRepo := repositories.NewProjectRepository(logger, testDb)

	getProjectQuery := queries.GetProjectQuery{ProjectId: 1}
	expectedErrorMessage := "Project not found!"

	// Act
	getResponse, err := testRepo.GetProject(getProjectQuery)

	// Assert
	assert.NotNil(t, getResponse)
	assert.NotNil(t, err)
	assert.EqualValues(t, expectedErrorMessage, err.Error())

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatal("Cleanup failed!: ", err.Error())
		}
	})
}

func Test_UpdateProject_ShouldReturnUpdatedProjectdto(t *testing.T) {
	// Arrange
	ctx := context.Background()
	logger := config.SetupLogger()
	secretClientConfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientConfig.GetSecretClient())
	postgresContainer := fixtures.SetupTestContainer(ctx, t)
	connectionString, err := postgresContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	testDb := dbConfig.GetDatabase(connectionString)
	testRepo := repositories.NewProjectRepository(logger, testDb)

	createProjectCommand := commands.CreateProjectCommand{
		Name:         "TestCreateProject",
		Description:  "Testdescription",
		Participants: "Participant",
		From:         time.Now(),
		To:           time.Now().Add(time.Hour),
	}

	// Act
	createResponse, err := testRepo.CreateProject(createProjectCommand)
	assert.NotEmpty(t, createResponse)
	assert.IsType(t, dtos.ProjectDto{}, createResponse)
	assert.NotNil(t, createResponse.ProjectId)
	assert.Nil(t, err)

	updateCommand := commands.UpdateProjectCommand{
		Name:         "UpdatedProjectName",
		Description:  "Project has been updated",
		Participants: "NewParticipant",
		From:         time.Now(),
		To:           time.Now().Add(time.Hour),
	}

	updateResponse, err := testRepo.UpdateProject(createResponse.ProjectId, updateCommand)

	assert.NotEmpty(t, updateResponse)
	assert.Nil(t, err)
	assert.IsType(t, dtos.ProjectDto{}, updateResponse)
	assert.NotEqualValues(t, createResponse.Name, updateResponse.Name)
	assert.NotEqualValues(t, createResponse.Description, updateResponse.Description)
	assert.NotEqualValues(t, createResponse.Participants, updateResponse.Participants)
	assert.NotEqualValues(t, createResponse.From, updateResponse.From)
	assert.NotEqualValues(t, createResponse.To, updateResponse.To)

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatal("Cleanup failed: ", err.Error())
		}
	})
}
