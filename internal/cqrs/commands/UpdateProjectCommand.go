package commands

import "time"

type UpdateProjectCommand struct {
	Name         string
	Description  string
	Participants string
	From         time.Time
	To           time.Time
}
