package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopRepository struct {
	db *mongo.Client
}

func (s ShopRepository) InsertItem(item model.DonatItem) error {
	db := s.db.Database(dbName)
	coll := db.Collection("shop")

	_, err := coll.InsertOne(context.Background(), item)
	return err
}

func NewShopRepository(db *mongo.Client) *ShopRepository {
	return &ShopRepository{db: db}
}
func (s ShopRepository) GetAll() ([]model.DonatItem, error) {
	db := s.db.Database(dbName)
	coll := db.Collection("shop")

	var result []model.DonatItem
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, err
}
