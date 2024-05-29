package utils

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func RunInTransaction(ctx context.Context, pool *pgxpool.Pool, txOptions pgx.TxOptions, f func(tx *sql.DB) error) error {
	tx, err := pool.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}
	defer handleTransaction(ctx, tx, err)

	db := stdlib.OpenDBFromPool(pool)

	return f(db)
}

func handleTransaction(ctx context.Context, tx pgx.Tx, err error) {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			zap.L().Error("failed to rollback transaction", zap.Error(rollbackErr))
		}
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			zap.L().Error("failed to commit transaction", zap.Error(commitErr))
		}
	}
}
