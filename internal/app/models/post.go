package models

import "time"

type Post struct {
	ID        int
	Content   string
	Title     string
	CreatedAt time.Time
}
