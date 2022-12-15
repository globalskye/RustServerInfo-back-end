package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type Service struct {
	UserI
	ClanI
}
type UserI interface {
	GetAllUsers() ([]model.User, error)
	GetOnline() ([]string, error)
}
type ClanI interface {
	GetAllClans() ([]model.Clan, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserI: NewUserService(repo.UserI),
		ClanI: NewClanService(repo.ClanI),
	}
}
