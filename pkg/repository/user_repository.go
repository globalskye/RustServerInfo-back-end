package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db *mongo.Client
}

func (u UserRepository) GetUserBySteamId(steamId int) (model.User, error) {
	db := u.db.Database("global")
	coll := db.Collection("users")

	var result model.User
	err := coll.FindOne(context.Background(), bson.D{{"steamId", steamId}}).Decode(&result)
	if err != nil {
		return model.User{}, err
	}
	return result, err
}

func (u UserRepository) GetUserByName(name string) (model.User, error) {
	db := u.db.Database("global")
	coll := db.Collection("users")

	var result model.User
	err := coll.FindOne(context.Background(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		return model.User{}, err
	}
	return result, err
}

func (u UserRepository) GetTopRaiders() ([]model.User, error) {
	db := u.db.Database("global")
	coll := db.Collection("users")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"raid", -1}})
	queryOptions.SetLimit(10)

	var result []model.User
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u UserRepository) GetTopKillers() ([]model.User, error) {
	db := u.db.Database("global")
	coll := db.Collection("users")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"killedPlayers", -1}})
	queryOptions.SetLimit(10)

	var result []model.User
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u UserRepository) GetUsers() ([]model.User, error) {
	db := u.db.Database("global")
	coll := db.Collection("users")
	var result []model.User
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u UserRepository) GetOnline() (model.Online, error) {
	db := u.db.Database("global")
	coll := db.Collection("online")
	var result model.Online
	err := coll.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		return model.Online{}, err
	}

	return result, err
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db: db}
}
