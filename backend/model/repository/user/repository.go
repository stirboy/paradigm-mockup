package user

import (
	"backend/app/database"
	"backend/model"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type UserRepository struct {
	Db *database.Database
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	})
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	user, err := pgxutil.SelectRow(ctx, tx, "SELECT * FROM users WHERE username = $1", []interface{}{username}, pgx.RowToAddrOfStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) InsertUser(ctx context.Context, user model.User) error {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	_, err = tx.Exec(ctx, "INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.Id, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
