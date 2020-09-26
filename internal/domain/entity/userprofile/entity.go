package userprofile

import "time"

type Entity struct {
	ID        int
	UserID    int
	UserName  string
	Thumbnail string
	Bio       string
	Gender    int
	Place     string
	Birth     *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type EntitySlice []*Entity
