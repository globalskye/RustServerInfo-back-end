package repository

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UserI
	ClanI
}

func NewRepository(m *mongo.Client) *Repository {
	return &Repository{
		UserI: NewUserRepository(m),
		ClanI: NewClanRepository(m),
	}
}

type UserI interface {
	GetUsers() ([]model.User, error)
	GetOnline() (model.Online, error)
}
type ClanI interface {
	GetClans() ([]model.Clan, error)
}
