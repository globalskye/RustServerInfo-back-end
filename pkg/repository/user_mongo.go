package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRepository struct {
	db *mongo.Client
}

func (u PlayerRepository) GetTopTime() ([]model.Player, error) {

	db := u.db.Database("global")
	coll := db.Collection("players")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"online", -1}})
	queryOptions.SetLimit(10)

	var result []model.Player
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err

}

func (u PlayerRepository) GetPlayerBySteamId(steamId int) (model.Player, error) {
	db := u.db.Database("global")
	coll := db.Collection("players")

	var result model.Player
	err := coll.FindOne(context.Background(), bson.D{{"steamId", steamId}}).Decode(&result)
	if err != nil {
		return model.Player{}, err
	}
	return result, err
}

func (u PlayerRepository) GetPlayerByName(name string) (model.Player, error) {
	db := u.db.Database("global")
	coll := db.Collection("players")

	var result model.Player
	err := coll.FindOne(context.Background(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		return model.Player{}, err
	}
	return result, err
}

func (u PlayerRepository) GetTopRaiders() ([]model.Player, error) {
	db := u.db.Database("global")
	coll := db.Collection("players")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"raid", -1}})
	queryOptions.SetLimit(10)

	var result []model.Player
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u PlayerRepository) GetTopKillers() ([]model.Player, error) {
	db := u.db.Database("global")
	coll := db.Collection("player")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"killedPlayers", -1}})
	queryOptions.SetLimit(10)

	var result []model.Player
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u PlayerRepository) GetPlayers() ([]model.Player, error) {
	db := u.db.Database("global")
	coll := db.Collection("player")
	var result []model.Player
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func (u PlayerRepository) GetOnline() (model.Online, error) {
	db := u.db.Database("global")
	coll := db.Collection("online")
	var result model.Online
	err := coll.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		return model.Online{}, err
	}

	return result, err
}

func NewPlayerRepository(db *mongo.Client) *PlayerRepository {
	return &PlayerRepository{db: db}
}
