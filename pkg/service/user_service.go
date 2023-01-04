package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type PlayerService struct {
	repo repository.PlayerI
}

func (u PlayerService) GetTopTime() ([]model.Player, error) {
	return u.repo.GetTopTime()
}

func (u PlayerService) GetPlayerBySteamId(steamId int) (model.Player, error) {
	return u.repo.GetPlayerBySteamId(steamId)
}

func (u PlayerService) GetPlayerByName(name string) (model.Player, error) {
	return u.repo.GetPlayerByName(name)
}

func (u PlayerService) GetTopRaiders() ([]model.Player, error) {
	return u.repo.GetTopRaiders()
}

func (u PlayerService) GetTopKillers() ([]model.Player, error) {
	return u.repo.GetTopKillers()
}

func (u PlayerService) GetPlayers() ([]model.Player, error) {
	return u.repo.GetPlayers()
}

func (u PlayerService) GetOnline() (model.Online, error) {
	return u.repo.GetOnline()
}

func NewPlayerService(repo repository.PlayerI) *PlayerService {
	return &PlayerService{repo: repo}
}
