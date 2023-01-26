package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Authorization
	PlayerI
	ClanI
	VkI
	UserI
	ShopI
}

type ShopI interface {
	GetAll() ([]model.DonatItem, error)
	InsertItem(item model.DonatItem) error
}

type UserI interface {
	GetUserById(id primitive.ObjectID) (model.User, error)
	CreateUser(user model.User) (interface{}, error)
	CheckUserName(name string) bool
	GetUserName(id primitive.ObjectID) (string, error)
}

type Authorization interface {
	GenerateAccessToken(username, password string) (string, error)
	ParseAccessToken(token string) (primitive.ObjectID, error)
}

type PlayerI interface {
	GetPlayers() ([]model.Player, error)
	GetOnline() (model.Online, error)
	GetTopKillers() ([]model.Player, error)
	GetTopRaiders() ([]model.Player, error)
	GetTopTime() ([]model.Player, error)
	GetPlayerByName(name string) (model.Player, error)
	GetPlayerBySteamId(steamId int) (model.Player, error)
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
		PlayerI:       NewPlayerService(repo.PlayerI),
		ClanI:         NewClanService(repo.ClanI),
		VkI:           NewVkService(repo.VkI),
		Authorization: NewAuthService(repo.UserI),
		UserI:         NewUserService(repo.UserI),
		ShopI:         NewShopService(repo.ShopI),
	}
}
