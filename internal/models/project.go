package models

import (
	"gin-backend/internal/dtos"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Participants string    `json:"participants"`
	IsActive     bool      `json:"is_active"`
	From         time.Time `json:"from" gorm:index:idx_from`
	To           time.Time `json:"to" gorm:index:idx_to`
	IsCompleted  bool      `json:"is_completed"`
	CreatedAt    time.Time `json:"created_at"`
}

func (p *Project) ToDto() dtos.ProjectDto {
	return dtos.ProjectDto{
		ProjectId:    p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Participants: p.Participants,
		IsActive:     p.IsActive,
		From:         p.From,
		To:           p.To,
		IsCompleted:  p.IsCompleted,
		CreatedAt:    p.CreatedAt,
	}
}
