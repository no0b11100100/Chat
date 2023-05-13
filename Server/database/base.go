package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Base struct {
	Collection *mongo.Collection
}

func (db *Base) Connect(client *mongo.Client, database string, collectionName string) {
	db.Collection = client.Database(database).Collection(collectionName)
}
