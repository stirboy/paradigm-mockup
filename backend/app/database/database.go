package database

import (
	"backend/app/config/requestid"
	"backend/app/database/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"strings"

	"go.uber.org/zap"

	"github.com/go-jet/jet/v2/generator/postgres"
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
	fixedSql := strings.ReplaceAll(removeWhitespaceAfterNewline(data.SQL), "\"", "")
	fields := []interface{}{
		"sql", fixedSql,
	}

	if data.Args != nil {
		fields = append(fields, "args", data.Args)
	}

	if reqID := requestid.GetReqID(ctx); reqID != uuid.Nil {
		fields = append(fields, "reqId", reqID)
	}

	tracer.log.Infow("Executing command", fields...)

	return ctx
}

func (tracer *myQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	//fields := []interface{}{
	//	"commandTag", data.CommandTag,
	//}
	//
	//if data.Err != nil {
	//	fields = append(fields, "err", data.Err)
	//}
	//
	//if reqID := requestid.GetReqID(ctx); reqID != uuid.Nil {
	//	fields = append(fields, "reqId", reqID)
	//}
	//tracer.log.Infow("Command completed", fields...)
}

func New(url string) (*Database, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		zap.L().Error("Unable to parse database config", zap.Error(err))
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
		zap.L().Error("Unable to connect to database", zap.Error(err))
		return nil, err
	}

	if err := migrateUp(context.Background(), dbPool); err != nil {
		zap.L().Error("Unable to migrate database", zap.Error(err))
		return nil, err
	}

	err = postgres.Generate("./.gen", postgres.DBConnection{
		Host:       config.ConnConfig.Host,
		Port:       int(config.ConnConfig.Port),
		User:       config.ConnConfig.User,
		Password:   config.ConnConfig.Password,
		DBName:     config.ConnConfig.Database,
		SchemaName: "public",
		SslMode:    "disable",
	})

	if err != nil {
		zap.L().Error("Unable to generate jet files", zap.Error(err))
		return nil, err
	}

	return &Database{
		DBPool: dbPool,
	}, nil
}

func migrateUp(ctx context.Context, db *pgxpool.Pool) error {
	err := utils.RunInTransaction(ctx, db, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func(tx *sql.DB) error {
		m, err := migrate.New("file://app/database/migration", db.Config().ConnString())
		if err != nil {
			zap.L().Error("Unable to create migration", zap.Error(err))
			return err
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			zap.L().Error("Unable to run migration", zap.Error(err))
			return err
		}

		return nil
	})

	return err
}

func removeWhitespaceAfterNewline(input string) string {
	var result strings.Builder
	skipSpaces := false

	for _, r := range input {
		if r == '\n' {
			skipSpaces = true
			result.WriteRune(' ')
		} else if skipSpaces && (r == ' ' || r == '\t') {
			continue
		} else {
			skipSpaces = false
			result.WriteRune(r)
		}
	}
	return result.String()
}
