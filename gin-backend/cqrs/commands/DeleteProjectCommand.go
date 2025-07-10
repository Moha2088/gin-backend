package commands

import "github.com/google/uuid"

type DeleteProjectCommand struct {
	ProjectId uuid.UUID
}
