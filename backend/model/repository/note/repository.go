package note

import (
	"backend/app/database"
	"backend/model"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
	"time"
)

type NoteRepository struct {
	Db *database.Database
}

func (r *NoteRepository) Insert(ctx context.Context, note model.Note) error {
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

	_, err = tx.Exec(ctx, "INSERT INTO notes (id, user_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", note.Id, note.UserId, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) GetByIdAndUserId(ctx context.Context, noteId uuid.UUID, userId uuid.UUID) (*model.Note, error) {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return nil, err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	note, err := pgxutil.SelectRow(ctx, tx, "Select * from notes where id = $1 and user_id = $2", []interface{}{noteId, userId}, pgx.RowToAddrOfStructByName[model.Note])
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *NoteRepository) GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]*model.Note, error) {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return nil, err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	notes, err := pgxutil.Select(ctx, tx, "Select * from notes where user_id = $1 order by created_at desc", []interface{}{userId}, pgx.RowToAddrOfStructByName[model.Note])
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) UpdateByUserId(ctx context.Context, id uuid.UUID, userId uuid.UUID, content string, title string) error {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	setValues := map[string]any{"updated_at": time.Now()}
	if content != "" {
		setValues["content"] = content
	}

	if title != "" {
		setValues["title"] = title
	}

	where := map[string]any{"user_id": userId, "id": id}

	err = pgxutil.UpdateRow(ctx, tx, "notes", setValues, where)
	if err != nil {
		return err
	}

	return nil
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]model.Note, error) {
	tx, err := r.Db.DBPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return nil, err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)
	rows, err := tx.Query(ctx, "SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
