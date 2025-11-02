package organization

import "time"

type Organization struct {
	ID        string
	Name      string
	ClientID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
