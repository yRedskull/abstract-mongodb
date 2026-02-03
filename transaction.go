package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func ExecTransactionFunctions(ctx context.Context, functions ...FunctionTransaction) error {
	db, err_db := GetDB()

	if err_db != nil {
		return err_db
	}

	session, err := db.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		for _, fn := range functions {
			if err := fn.Function(sessCtx, fn.Params...); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}
