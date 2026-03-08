package handlers

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Row abstracts a single result row for testability.
type Row interface {
	Scan(dest ...any) error
}

// Rows abstracts a result set for testability.
type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close()
	Err() error
}

// DB is the database interface used by handlers.
type DB interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Ping(ctx context.Context) error
}

// PgxPool adapts *pgxpool.Pool to implement the DB interface.
type PgxPool struct {
	*pgxpool.Pool
}

func (p *PgxPool) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return p.Pool.Query(ctx, sql, args...)
}

func (p *PgxPool) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return p.Pool.QueryRow(ctx, sql, args...)
}
