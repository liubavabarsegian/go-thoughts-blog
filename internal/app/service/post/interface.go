package post

import "day06/internal/app/models"

// Service определяет методы, которые должен реализовывать сервис постов
type Service interface {
	GetAllPosts(page int) ([]models.Post, error)
	// GetPostByID(id int) (*models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	// UpdatePost(id int, post *models.Post) error
	// DeletePost(id int) error
}
