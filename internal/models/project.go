package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ProjectId    uuid.UUID `"json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Participants []string  `json:"participants"`
	IsActive     bool      `json:"is_active"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	IsCompleted  bool      `json:"is_completed"`
	CreatedAt    time.Time `json:"created_at"`
}
