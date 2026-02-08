package handlers

import "github.com/jackc/pgx/v5/pgxpool"

// H holds shared dependencies injected into all handlers.
type H struct {
	DB *pgxpool.Pool
}
