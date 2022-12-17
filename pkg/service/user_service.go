package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type UserService struct {
	repo repository.UserI
}

func (u UserService) GetUserBySteamId(steamId int) (model.User, error) {
	return u.repo.GetUserBySteamId(steamId)
}

func (u UserService) GetUserByName(name string) (model.User, error) {
	return u.repo.GetUserByName(name)
}

func (u UserService) GetTopRaiders() ([]model.User, error) {
	return u.repo.GetTopRaiders()
}

func (u UserService) GetTopKillers() ([]model.User, error) {
	return u.repo.GetTopKillers()
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
