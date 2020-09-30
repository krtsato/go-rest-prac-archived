package user

import "time"

type Entity struct {
	ID        int
	ProfileID int
	AuthID    string
	Mail      string
	Phone     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type EntitySlice []*Entity
