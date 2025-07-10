package queries

import "github.com/google/uuid"

type GetProjectQuery struct {
	ProjectId uuid.UUID
}
