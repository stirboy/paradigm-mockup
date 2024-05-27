package note

import (
	"backend/app/database"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type NoteRepository struct {
	Db *database.Database
}

func (r *NoteRepository) Insert(ctx context.Context, note Note) error {
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

	_, err = pgxutil.InsertRowReturning(ctx, tx, "notes", map[string]any{"id": note.Id, "user_id": note.UserId, "title": note.Title, "content": note.Content, "is_archived": note.IsArchived, "parent_id": note.ParentId, "cover_image": note.CoverImage, "icon": note.Icon, "created_at": note.CreatedAt, "updated_at": note.UpdatedAt}, "id", pgx.RowTo[uuid.UUID])
	if err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) GetByIdAndUserId(ctx context.Context, noteId uuid.UUID, userId uuid.UUID) (*Note, error) {
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

	note, err := pgxutil.SelectRow(ctx, tx, "Select * from notes where id = $1 and user_id = $2", []interface{}{noteId, userId}, pgx.RowToAddrOfStructByName[Note])
	if err != nil {
		return nil, err
	}

	return note, nil
}

type Options struct {
	ParentId *uuid.UUID
}

func (r *NoteRepository) GetAllByUserId(ctx context.Context, userId uuid.UUID, options *Options) ([]*Note, error) {
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

	var query string
	var values []interface{}

	if options.ParentId != nil {
		query = "Select * from notes where user_id = $1 and parent_id = $2 order by created_at desc"
		values = []interface{}{userId, options.ParentId}
	} else {
		query = "Select * from notes where user_id = $1 and parent_id is null order by created_at desc"
		values = []interface{}{userId}
	}

	notes, err := pgxutil.Select(ctx, tx, query, values, pgx.RowToAddrOfStructByName[Note])
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

func (r *NoteRepository) GetAll(ctx context.Context) ([]Note, error) {
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

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.Id,
			&note.UserId,
			&note.Title,
			&note.Content,
			&note.IsArchived,
			&note.ParentId,
			&note.CoverImage,
			&note.Icon,
			&note.CreatedAt,
			&note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
