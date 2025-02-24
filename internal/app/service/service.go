package service

import (
	"day06/internal/app/repository"
	"day06/internal/app/service/post"
)

type Service struct {
	Post post.Service
	// User    user.Service
	// Comment comment.Service
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post: post.NewService(*repo.Post),
		// User:    user.NewService(repo.User),
		// Comment: comment.NewService(repo.Comment),
	}
}
