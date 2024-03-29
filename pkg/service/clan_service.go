package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type ClanService struct {
	repo repository.ClanI
}

func (c ClanService) GetClanByName(name string) (model.Clan, error) {
	return c.repo.GetClanByName(name)
}

func (c ClanService) GetTopClans() ([]model.Clan, error) {
	return c.repo.GetTopClans()
}

func (c ClanService) GetClans() ([]model.Clan, error) {

	return c.repo.GetClans()
}

func NewClanService(repo repository.ClanI) *ClanService {
	return &ClanService{repo: repo}
}
