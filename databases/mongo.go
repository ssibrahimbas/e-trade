package databases

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	c   *mongo.Client
	db  *mongo.Database
	ctx context.Context
}

func NewMongoDB(uri string, dbName string) *MongoDB {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db := client.Database(dbName)
	return &MongoDB{
		c:   client,
		ctx: ctx,
		db:  db,
	}
}

func (m *MongoDB) GetCollection(n string) *mongo.Collection {
	return m.db.Collection(n)
}
