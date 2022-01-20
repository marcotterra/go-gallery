package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MongoClientTimeout = 5

type FolderRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewFolderRepository(connection string) (*FolderRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &FolderRepository{
		client:     client,
		database:   client.Database("gallery"),
		collection: client.Database("gallery").Collection("folders"),
	}, nil
}

func (repository *FolderRepository) Store(title string) error {
	// result, err := repository.collection.InsertOne(bson.D{
	// 	{Key: "Test"},
	// })
	return nil
}
