package database

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	DBPool *pgxpool.Pool
}

type myQueryTracer struct {
	log *zap.SugaredLogger
}

func (tracer *myQueryTracer) TraceQueryStart(
	ctx context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryStartData) context.Context {
	tracer.log.Infow("Executing command", "sql", data.SQL, "args", data.Args)

	return ctx
}

func (tracer *myQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	tracer.log.Infow("Command completed", "commandTag", data.CommandTag, "err", data.Err)
}

func New(url string) (*Database, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse database config: %v\n", err)
		return nil, err
	}

	config.ConnConfig.Tracer = &myQueryTracer{
		log: zap.S(),
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err := migrateUp(context.Background(), dbPool); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to migrate database: %v\n", err)
		return nil, err
	}

	return &Database{
		DBPool: dbPool,
	}, nil
}

func migrateUp(ctx context.Context, db *pgxpool.Pool) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start transaction: %v\n", err)
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	m, err := migrate.New("file://app/database/migration", db.Config().ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create migration: %v\n", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Fprintf(os.Stderr, "Unable to run migration: %v\n", err)
		return err
	}

	return nil
}
