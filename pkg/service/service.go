package service

import "github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
