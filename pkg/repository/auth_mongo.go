package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type AuthRepository struct {
	db *mongo.Client
}

func NewAuthRepository(db *mongo.Client) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) GetUserByName(name string) (model.User, error) {
	db := a.db.Database("global")
	coll := db.Collection("users")

	var result model.User
	err := coll.FindOne(context.Background(), bson.D{{"username", name}}).Decode(&result)
	if err != nil {
		return model.User{}, err
	}
	return result, err
}

func (a *AuthRepository) GetUserById(id primitive.ObjectID) (model.User, error) {
	db := a.db.Database("global")
	coll := db.Collection("users")

	var result model.User
	err := coll.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		return model.User{}, err
	}
	return result, err
}

func (a *AuthRepository) CreateUser(user model.User) (interface{}, error) {
	db := a.db.Database(os.Getenv("MONGODB_DATABASE"))
	collection := db.Collection("users")

	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID, err
}
func (a *AuthRepository) GetUser(username, password string) (model.User, error) {
	db := a.db.Database(os.Getenv("MONGODB_DATABASE"))
	collection := db.Collection("users")

	var user model.User
	err := collection.FindOne(context.Background(), bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, err

}
