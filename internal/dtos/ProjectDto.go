package dtos

import (
	"time"
)

type ProjectDto struct {
	ProjectId    uint
	Name         string
	Description  string
	Participants []string
	IsActive     bool
	From         time.Time
	To           time.Time
	IsCompleted  bool
	CreatedAt    time.Time
}
