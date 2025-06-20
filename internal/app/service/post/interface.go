package post

import "day06/internal/app/models"

// Service определяет методы, которые должен реализовывать сервис постов
type Service interface {
	GetAllPosts(page uint) ([]models.Post, error)
	GetPost(id uint) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
}
