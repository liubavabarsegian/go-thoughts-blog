package service

import (
	"day06/internal/app/repository"
	"day06/internal/app/service/login"
	"day06/internal/app/service/post"
)

type Service struct {
	Post  post.Service
	Login login.Service
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post:  post.NewService(*repo.Post),
		Login: login.NewService(),
	}
}
