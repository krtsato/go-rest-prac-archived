package entity

import "time"

type Userprofile struct {
	UserID    int
	UserName  string
	Thumbnail string
	Bio       string
	Gender    int
	Mail      string
	Phone     string
	Place     string
	Birth     *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
