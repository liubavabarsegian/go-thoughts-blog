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

// GetPost находит пост по ID
func (s *PostService) GetPost(id uint) (models.Post, error) {
	return s.repo.GetPost(id)
}

// GetAllPosts возвращает все посты
func (s *PostService) GetAllPosts(page uint) ([]models.Post, error) {
	return s.repo.GetAllPosts(page)
}

// CreatePost — обработчик создания поста
func (s *PostService) CreatePost(postModel models.Post) (models.Post, error) {
	return s.repo.CreatePost(postModel)
}
