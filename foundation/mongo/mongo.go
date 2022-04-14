package mongo

import (
	"context"
	"errors"
	"ginson/foundation/cfg"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	mg *MongoInstance
)

func InitMongo(cfg *cfg.AppConfig) error {
	if cfg.MongoConfig == nil {
		return nil
	}
	var err error
	mg, err = connectMongo(cfg)
	if err != nil {
		return err
	}
	return nil
}

func Mongo() *MongoInstance {
	if mg == nil {
		panic(errors.New("mongodb is not ready"))
	}
	return mg
}

func connectMongo(cfg *cfg.AppConfig) (*MongoInstance, error) {
	option := options.Client().ApplyURI(cfg.MongoConfig.Uri).
		SetConnectTimeout(time.Duration(cfg.MongoConfig.ConnectTimeout) * time.Second).
		SetMaxConnecting(cfg.MongoConfig.MaxOpenConn).
		SetMaxPoolSize(cfg.MongoConfig.MaxPoolSize).SetMinPoolSize(cfg.MongoConfig.MinPoolSize)
	client, err := mongo.NewClient(option)

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &MongoInstance{
		Client: client,
		Db:     client.Database(cfg.MongoConfig.Db),
	}, nil
}
