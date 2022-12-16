package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClanRepository struct {
	db *mongo.Client
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
