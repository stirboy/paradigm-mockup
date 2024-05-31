package utils

import (
	"backend/app/database/dberrors"
	"context"
	"database/sql"
	"errors"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

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

func Query(ctx context.Context, q postgres.SelectStatement, db qrm.Queryable, destination interface{}) error {
	err := q.QueryContext(ctx, db, destination)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return dberrors.NotFoundError
		}
		return err
	}

	return nil
}

func Exec(ctx context.Context, q postgres.Statement, db qrm.Executable) error {
	res, err := q.ExecContext(ctx, db)
	if err != nil {

	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return dberrors.OptimisticLockingError
	}

	return err
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
