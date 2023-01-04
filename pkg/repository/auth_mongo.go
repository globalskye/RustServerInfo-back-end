package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type AuthRepository struct {
	db *mongo.Client
}

func (a *AuthRepository) GetUserById(id int) ([]model.User, error) {
	panic("qwe")
}

func NewAuthRepository(db *mongo.Client) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) CreateUser(user model.User) (int, error) {
	db := a.db.Database(os.Getenv("MONGODB_DATABASE"))
	collection := db.Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	collection.InsertOne(context.Background(), user)

	return 0, err
}
func (a *AuthRepository) GetUser(username, password string) (model.User, error) {
	panic("qwe")

}
