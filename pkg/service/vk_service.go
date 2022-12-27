package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
)

type VkService struct {
	repo repository.VkI
}

func (v VkService) GetVkPosts() (model.VkPost, error) {
	return v.repo.GetVkPosts()
}

func NewVkService(repo repository.VkI) *VkService {
	return &VkService{repo: repo}
}
