package repository

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	PlayerI
	Authorization
	ClanI
	VkI
}

func NewRepository(m *mongo.Client) *Repository {
	return &Repository{
		PlayerI:       NewPlayerRepository(m),
		ClanI:         NewClanRepository(m),
		VkI:           NewVkRepository(m),
		Authorization: NewAuthRepository(m),
	}
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
	GetUserById(id int) ([]model.User, error)
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
