package unittests

import (
	"context"
	"gin-backend/internal/config"
	"gin-backend/internal/cqrs/commands"
	"gin-backend/internal/dtos"
	"gin-backend/internal/repositories"
	"gin-backend/tests/fixtures"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testRepository struct {
}

func TestProjectRepository(t *testing.T) {
	ctx := context.Background()
	logger := config.SetupLogger()
	secretClientConfig := config.NewSecretClientConfig(logger)
	dbConfig := config.NewDatabaseConfig(logger, secretClientConfig.GetSecretClient())
	postgresContainer := fixtures.SetupTestContainer(ctx, t)
	connectionString, err := postgresContainer.ConnectionString(ctx)
	assert.NoError(t, err)
	testDb := dbConfig.GetDatabase(connectionString)

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatal("Cleanup failed: ", err)
		}
	})

	// CreateProject_ShouldReturnCreatedProjectdto_WhenCreatingProject
	createProjectCommand := commands.CreateProjectCommand{
		Name:         "TestProject",
		Description:  "Testdescription",
		Participants: "Participant",
		From:         time.Now(),
		To:           time.Now().Add(time.Hour),
	}

	testRepo := repositories.NewProjectRepository(logger, testDb)
	response := testRepo.CreateProject(createProjectCommand)

	assert.NotEmpty(t, response)
	assert.IsType(t, dtos.ProjectDto{}, response)
	assert.NotNil(t, response.ProjectId)
}
