package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindParams struct {
	Collection string
	Filter     bson.M
	Fields     []string
}

type CountDocumentsParams struct {
	Collection string
	Filter any
}

type FunctionTransaction struct {
	Function func (sessCtx mongo.SessionContext, params ...any) error
	Params []any
}

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}
