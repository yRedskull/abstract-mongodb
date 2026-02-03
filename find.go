package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(ctx context.Context, obj any, findParams FindParams, findOneOptions ...*options.FindOneOptions) error {
	if err_val := ValFindParams(findParams); err_val != nil {
		return err_val
	}

	db, err_db := GetDB()

	if err_db != nil {
		return err_db
	}

	col := db.Database.Collection(findParams.Collection)

	opts := BuildFindOneOptions(findParams, findOneOptions...)

	err := col.FindOne(ctx, findParams.Filter, opts).Decode(obj)
	if err != nil {
		return err
	}

	return nil
}

func Find[T any](ctx context.Context, list *[]T, findParams FindParams, findOptions ...*options.FindOptions) error {
	db, err_db := GetDB()

	if err_db != nil {
		return err_db
	}

	col := db.Database.Collection(findParams.Collection)

	opts := BuildFindOptions(findParams, findOptions...)

	cursor, err := col.Find(ctx, findParams.Filter, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, list); err != nil {
		return err
	}

	if len(*list) == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func FindAggregate(ctx context.Context, collection string, filter bson.M, projection bson.M) ([]bson.M, error) {
	db, _err := GetDB()

	if _err != nil {
		return nil, _err
	}

	col := db.Database.Collection(collection)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$project", Value: projection}},
	}

	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return results, nil
}
