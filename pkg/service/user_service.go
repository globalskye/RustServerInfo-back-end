package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type UserService struct {
	repo repository.UserI
}

func (u UserService) GetUsers() ([]model.User, error) {
	return u.repo.GetUsers()
}

func (u UserService) GetOnline() (model.Online, error) {
	return u.repo.GetOnline()
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{repo: repo}
}
