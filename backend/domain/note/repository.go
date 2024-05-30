package note

import (
	"backend/app/database"
	"backend/app/database/utils"
	"context"
	"database/sql"
	"time"

	"backend/domain/note/model"
	. "backend/domain/note/table"

	"go.uber.org/zap"

	. "github.com/go-jet/jet/v2/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Db *database.Database
}

func (r *Repository) Insert(ctx context.Context, note model.Notes) (*uuid.UUID, error) {
	var dest struct {
		ID uuid.UUID `alias:"notes.id"`
	}
	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func(tx *sql.DB) error {
		q := Notes.INSERT(Notes.AllColumns).MODEL(note).RETURNING(Notes.ID)
		return q.QueryContext(ctx, tx, &dest)
	})

	if err != nil {
		zap.L().Error("cant insert note", zap.Error(err))
		return nil, err
	}

	return &dest.ID, nil
}

func (r *Repository) GetByIdAndUserId(ctx context.Context, noteId uuid.UUID, userId uuid.UUID) (*model.Notes, error) {

	note := &model.Notes{}
	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	}, func(tx *sql.DB) error {
		q := SELECT(Notes.AllColumns).FROM(Notes).WHERE(Notes.ID.EQ(UUID(noteId)).AND(Notes.UserID.EQ(UUID(userId))))
		return q.QueryContext(ctx, tx, note)
	})

	if err != nil {
		zap.L().Error("cant get note", zap.Error(err))
		return nil, err
	}

	return note, nil
}

type Options struct {
	ParentId     *uuid.UUID
	ArchivedOnly bool
}

func (r *Repository) GetAllByUserId(ctx context.Context, userId uuid.UUID, options *Options) ([]*model.Notes, error) {

	var notes []*model.Notes

	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadOnly,
	}, func(tx *sql.DB) error {
		condition := Notes.UserID.EQ(UUID(userId)).AND(Notes.IsArchived.EQ(Bool(options.ArchivedOnly)))

		if options.ParentId != nil {
			condition = condition.AND(Notes.ParentID.EQ(UUID(options.ParentId)))
		} else {
			condition = condition.AND(Notes.ParentID.IS_NULL())
		}
		q := SELECT(Notes.AllColumns).FROM(Notes).WHERE(condition).ORDER_BY(Notes.CreatedAt.DESC())

		return q.QueryContext(ctx, tx, &notes)
	})

	if err != nil {
		zap.L().Error("cant select notes", zap.Error(err))
		return nil, err
	}

	return notes, nil
}

func (r *Repository) ArchiveNotes(ctx context.Context, userId uuid.UUID, parentId uuid.UUID) error {
	return r.archivedNotesRecursive(ctx, userId, parentId, true)
}

func (r *Repository) RestoreNotes(ctx context.Context, userId uuid.UUID, parentId uuid.UUID) error {
	return r.archivedNotesRecursive(ctx, userId, parentId, false)
}

func (r *Repository) archivedNotesRecursive(ctx context.Context, userId uuid.UUID, parentId uuid.UUID, shouldArchive bool) error {
	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func(tx *sql.DB) error {
		children := CTE("children")
		q := WITH_RECURSIVE(
			children.AS(
				SELECT(Notes.ID, Notes.ParentID).FROM(Notes).WHERE(Notes.ID.EQ(UUID(parentId)).AND(Notes.UserID.EQ(UUID(userId)))).UNION(
					SELECT(Notes.ID, Notes.ParentID).FROM(Notes.INNER_JOIN(children, Notes.ID.From(children).EQ(Notes.ParentID))),
				),
			),
		)(Notes.UPDATE(Notes.IsArchived).SET(Bool(shouldArchive)).WHERE(Notes.ID.IN(SELECT(Notes.ID.From(children)).FROM(children))))

		_, err := q.ExecContext(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		zap.L().Error("cant archive notes", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repository) UpdateByUserId(ctx context.Context, id uuid.UUID, userId uuid.UUID, content string, title string) error {

	err := utils.RunInTransaction(ctx, r.Db.DBPool, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func(tx *sql.DB) error {
		n := model.Notes{
			Title:     title,
			Content:   &content,
			UpdatedAt: time.Now(),
		}
		q := Notes.UPDATE(Notes.Title, Notes.Content).MODEL(n).WHERE(Notes.ID.EQ(UUID(id)).AND(Notes.UserID.EQ(UUID(userId))))
		_, err := q.ExecContext(ctx, tx)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		zap.L().Error("cant update notes", zap.Error(err))
		return err
	}

	return nil
}
