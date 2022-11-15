package service

import (
	"context"
	"fmt"
	"mongodb-trx/shared/model/repository"
)

// WithoutTransaction is helper function that simplify the readonly db
func WithoutTransaction[T any](ctx context.Context, trx repository.WithoutTransactionDB, trxFunc func(dbCtx context.Context) (*T, error)) (*T, error) {
	dbCtx, err := trx.GetDatabase(ctx)
	if err != nil {
		return nil, err
	}

	defer func(trx repository.WithoutTransactionDB, ctx context.Context) {
		err := trx.Close(ctx)
		if err != nil {
			return
		}
	}(trx, dbCtx)

	return trxFunc(dbCtx)
}

// WithTransaction is helper function that simplify the transaction execution handling
func WithTransaction[T any](ctx context.Context, trx repository.WithTransactionDB, trxFunc func(dbCtx context.Context) (*T, error)) (*T, error) {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf(">>>>> Rollback panic\n")
			err = trx.RollbackTransaction(dbCtx)
			panic(p)

		} else if err != nil {
			fmt.Printf(">>>>> Rollback error\n")
			err = trx.RollbackTransaction(dbCtx)

		} else {
			fmt.Printf(">>>>> Commit normal\n")
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	var t *T
	t, err = trxFunc(dbCtx)
	if err != nil {
		return nil, err
	}

	return t, err
}
