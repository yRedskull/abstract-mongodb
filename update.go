package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	db, err_db := GetDB()

	if err_db != nil {
		return nil, err_db
	}

	res, err := db.Database.Collection(collection).UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateMany(ctx context.Context, collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	db, err_db := GetDB()

	if err_db != nil {
		return nil, err_db
	}

	res, err := db.Database.Collection(collection).UpdateMany(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return res, nil
}
