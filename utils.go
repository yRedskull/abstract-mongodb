package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildFindOneOptions(findParams FindParams, findOptions ...*options.FindOneOptions) *options.FindOneOptions {
	var opts *options.FindOneOptions

	if len(findOptions) > 0 {
		opts = findOptions[0]
	} else {
		opts = options.FindOne()
	}

	if len(findParams.Fields) > 0 {
		projection := bson.M{}
		for _, field := range findParams.Fields {
			projection[field] = 1
		}
		opts.SetProjection(projection)
	}

	return opts
}

func BuildFindOptions(findParams FindParams, findOptions ...*options.FindOptions) *options.FindOptions {
	var opts *options.FindOptions

	if len(findOptions) > 0 && findOptions[0] != nil {
		opts = findOptions[0]
	} else {
		opts = options.Find()
	}

	if len(findParams.Fields) > 0 {
		projection := bson.M{}
		for _, field := range findParams.Fields {
			projection[field] = 1
		}
		opts.SetProjection(projection)
	}

	return opts
}


func EnsureIndex(ctx context.Context, collection string, indexModel mongo.IndexModel) error {
	var ok_index bool

	modelKeys, ok_index := indexModel.Keys.(bson.D)

	if !ok_index {
		return errors.New("indexModel não é do tipo []bson.D")
	}

	db, err := GetDB()
	if err != nil {
		return fmt.Errorf("falha ao conectar ao MongoDB: %w", err)
	}

	col := db.Database.Collection(collection)

	cursor, err := col.Indexes().List(ctx)
	if err != nil {
		return fmt.Errorf("falha ao listar índices: %w", err)
	}

	var indexes []bson.M
	if err = cursor.All(ctx, &indexes); err != nil {
		return fmt.Errorf("falha ao decodificar índices: %w", err)
	}

	for _, index := range indexes {
		if key, ok := index["key"].(bson.M); ok {
			for _, index := range modelKeys {
				if key[index.Key] == 1 {
					return nil
				}
			}
		}
	}

	_, err = col.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("erro ao criar índice %s: %w", indexModel.Keys, err)
	}

	return nil
}
