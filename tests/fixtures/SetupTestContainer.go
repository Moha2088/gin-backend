package fixtures

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestContainer(ctx context.Context, t *testing.T) *postgres.PostgresContainer {
	postgresContainer, err := postgres.Run(ctx, "postgres:latest",
		postgres.WithDatabase("unit_test_db"), testcontainers.WithWaitStrategy(wait.ForListeningPort("5432/tcp")))

	if err != nil {
		t.Fatal(err)
	}

	return postgresContainer
}
