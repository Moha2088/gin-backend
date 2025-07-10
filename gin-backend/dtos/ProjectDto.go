package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ProjectDto struct {
	ProjectId    uuid.UUID
	Name         string
	Description  string
	Participants []string
	IsActive     bool
	From         time.Time
	To           time.Time
	IsCompleted  bool
	CreatedAt    time.Time
}
