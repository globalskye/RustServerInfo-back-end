package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type Service struct {
	UserI
	ClanI
	VkI
}
type UserI interface {
	GetUsers() ([]model.User, error)
	GetOnline() (model.Online, error)
	GetTopKillers() ([]model.User, error)
	GetTopRaiders() ([]model.User, error)
	GetTopTime() ([]model.User, error)
	GetUserByName(name string) (model.User, error)
	GetUserBySteamId(steamId int) (model.User, error)
}
type ClanI interface {
	GetClans() ([]model.Clan, error)
	GetTopClans() ([]model.Clan, error)
	GetClanByName(name string) (model.Clan, error)
}
type VkI interface {
	GetVkPosts() (model.VkPost, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserI: NewUserService(repo.UserI),
		ClanI: NewClanService(repo.ClanI),
		VkI:   NewVkService(repo.VkI),
	}
}
