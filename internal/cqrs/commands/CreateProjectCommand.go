package commands

import "time"

type CreateProjectCommand struct {
	Name         string
	Description  string
	Participants string
	From         time.Time
	To           time.Time
}
