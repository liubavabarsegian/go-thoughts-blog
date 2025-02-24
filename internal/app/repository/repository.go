package repository

import (
	"database/sql"
)

type Repository struct {
	Post *PostRepository
	// User    UserRepository
	// Comment CommentRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post: NewPostRepository(db),
		// User:    NewUserRepository(db),
		// Comment: NewCommentRepository(db),
	}
}
