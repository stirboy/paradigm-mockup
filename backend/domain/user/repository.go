package user

import (
	"backend/app/database"
	"backend/app/database/utils"
	"context"
	"database/sql"

	"backend/domain/user/model"
	. "backend/domain/user/table"

	"go.uber.org/zap"

	. "github.com/go-jet/jet/v2/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Db *database.Database
}

func (r *Repository) GetUserById(ctx context.Context, id uuid.UUID) (*model.Users, error) {

	user := &model.Users{}

	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	}, func(tx *sql.DB) error {
		q := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.ID.EQ(UUID(id)))
		return utils.Query(ctx, q, tx, user)
	})

	if err != nil {
		zap.L().Error("cant get user by id", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*model.Users, error) {

	var user = &model.Users{}

	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	}, func(tx *sql.DB) error {
		q := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.Username.EQ(String(username)))
		return utils.Query(ctx, q, tx, user)
	})

	if err != nil {
		zap.L().Error("cant get user by username", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (r *Repository) InsertUser(ctx context.Context, user model.Users) error {
	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func(tx *sql.DB) error {
		q := Users.INSERT(Users.AllColumns).MODEL(user)
		_, err := q.ExecContext(ctx, tx)
		return err
	})

	if err != nil {
		zap.L().Error("cant insert user", zap.Error(err))
		return err
	}
	return nil
}
