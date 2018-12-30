package store

import (
	"context"
	"fmt"
	"github.com/axle-h/cheese/config"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type MongoContext struct {
	config config.MongoConfig
	client *mongo.Client
}

func NewMongoContext(mongoConfig config.MongoConfig) (MongoContext, error) {
	client, err := mongo.NewClient(mongoConfig.ConnectionString)
	mongoContext := MongoContext {mongoConfig, client}

	if err != nil {
		return mongoContext, err
	}

	err = mongoContext.client.Connect(mongoContext.timeout())

	if err != nil {
		return mongoContext, err
	}

	// Check the connection
	err = client.Ping(mongoContext.timeout(), nil)

	if err != nil {
		return mongoContext, fmt.Errorf("cannot connect to mongodb: %v", err)
	}

	return mongoContext, err
}

func (context MongoContext) GetCollection(name string) *mongo.Collection {
	return context.client.Database(context.config.Database).Collection(name)
}

func (_ MongoContext) timeout() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
