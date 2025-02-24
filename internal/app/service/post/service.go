package post

import (
	"day06/internal/app/models"
	repository "day06/internal/app/repository"
)

// postService реализует интерфейс Service
type PostService struct {
	repo repository.PostRepository
}

// NewPostService создает новый сервис для постов
func NewService(repo repository.PostRepository) Service {
	return &PostService{repo: repo}
}

// GetAllPosts возвращает все посты
func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}
