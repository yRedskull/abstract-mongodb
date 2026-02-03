package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountDocuments(ctx context.Context, countParams CountDocumentsParams, opts ...*options.CountOptions) (int64, error){
	db, err_db := GetDB()

	if err_db != nil {
		return 0, err_db
	}

	count, err := db.Database.Collection(countParams.Collection).CountDocuments(ctx, countParams.Filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}