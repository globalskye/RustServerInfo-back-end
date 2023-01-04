package repository

import (
	"context"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VkRepository struct {
	db *mongo.Client
}

func (v VkRepository) GetVkPosts() (model.VkPost, error) {
	db := v.db.Database("global")
	coll := db.Collection("vk")
	var result model.VkPost
	err := coll.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		return model.VkPost{}, err
	}

	return result, err

}

func NewVkRepository(db *mongo.Client) *VkRepository {
	return &VkRepository{db: db}
}
