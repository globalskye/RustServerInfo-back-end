package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type ShopService struct {
	repo repository.ShopI
}

func (s ShopService) InsertItem(item model.DonatItem) error {
	return s.repo.InsertItem(item)
}

func NewShopService(repo repository.ShopI) *ShopService {
	return &ShopService{repo: repo}
}

func (s ShopService) GetAll() ([]model.DonatItem, error) {
	return s.repo.GetAll()
}
