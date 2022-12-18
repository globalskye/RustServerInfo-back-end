package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClanRepository struct {
	db *mongo.Client
}

func (c ClanRepository) GetClanByName(name string) (model.Clan, error) {
	db := c.db.Database("global")
	coll := db.Collection("clans")

	var result model.Clan
	err := coll.FindOne(context.Background(), bson.D{{"name", name}}).Decode(&result)
	if err != nil {
		return model.Clan{}, err
	}
	return result, err
}

func (c ClanRepository) GetTopClans() ([]model.Clan, error) {
	db := c.db.Database("global")
	coll := db.Collection("clans")

	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{"level", -1}})
	queryOptions.SetLimit(10)

	var result []model.Clan
	cursor, err := coll.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	return result, err
}

func (c ClanRepository) GetClans() ([]model.Clan, error) {
	db := c.db.Database("global")
	coll := db.Collection("clans")
	var result []model.Clan
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}

func NewClanRepository(db *mongo.Client) *ClanRepository {
	return &ClanRepository{db: db}
}
