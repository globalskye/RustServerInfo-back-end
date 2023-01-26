package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repo repository.UserI
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{repo: repo}
}
func (a *UserService) GetUserName(id primitive.ObjectID) (string, error) {
	user, err := a.repo.GetUserById(id)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (a *UserService) CheckUserName(name string) bool {
	_, err := a.repo.GetUserByName(name)

	if err != mongo.ErrNoDocuments {
		return true
	}

	return false
}
func (a *UserService) GetUserById(id primitive.ObjectID) (model.User, error) {
	return a.repo.GetUserById(id)
}
func (a *UserService) CreateUser(user model.User) (interface{}, error) {
	user.Password = generateHashPassword(user.Password)
	user.Id = primitive.NewObjectID()
	return a.repo.CreateUser(user)

}
