package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type Service struct {
	UserI
}
type UserI interface {
	GetAllUsers() ([]model.User, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserI: NewUserService(repo.UserI),
	}
}
