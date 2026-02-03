package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(ctx context.Context, colletion string, new_document any) (*mongo.InsertOneResult, error) {
	db, err_db := GetDB()

	if err_db != nil {
		return nil, err_db
	}

	result, err_insert_document := db.Database.Collection(colletion).InsertOne(ctx, new_document)

	if err_insert_document != nil {
		return nil, err_insert_document
	}

	return result, nil
}
