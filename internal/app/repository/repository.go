package repository

import (
	"database/sql"
)

type Repository struct {
	Post *PostRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post: NewPostRepository(db),
	}
}
