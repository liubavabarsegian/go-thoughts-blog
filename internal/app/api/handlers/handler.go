package handler

import (
	"database/sql"
	"day06/internal/app/models"
	repo "day06/internal/app/repositories/post"
)

func GetAllPosts(db *sql.DB) []models.Post {
	result := repo.GetAll(db)
	return result
}
