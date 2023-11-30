package mongodb

import (
	"context"

	"github.com/yagoinacio/portfolio-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func OpenConnection() error {
	uri := config.GetDBConfig().ConnURI

	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = conn.Ping(context.TODO(), readpref.Primary())

	client = conn

	return err
}

func GetCollection(name string) *mongo.Collection {
	db := config.GetDBConfig().Database
	return client.Database(db).Collection(name)
}
