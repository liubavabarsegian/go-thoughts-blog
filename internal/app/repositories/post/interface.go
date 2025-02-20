package post

import "day06/internal/app/models"

type PostRepository interface {
	GetByID(id int) (*models.Post, error)
	GetAll() ([]*models.Post, error)
	Create(post *models.Post) error
}
