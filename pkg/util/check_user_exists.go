package util

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsUserExits(query *bson.D, db *mongo.Database) (bool, error) {
	var value, err = db.Collection("User").CountDocuments(context.Background(), query)
	return value == 1, err
}
