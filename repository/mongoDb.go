package repository

import (
	"context"
	"go-ms-profile/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	ctx            context.Context
	userCollection *mongo.Collection
	db             *mongo.Client
}


func NewMongoRepository(db *mongo.Client, cfg *config.ServerConfig) *MongoRepository {
	return &MongoRepository{
		ctx:            context.Background(),
		userCollection: db.Database(cfg.MongoDb.Database).Collection("profile"),
		db:             db,
	}
}

