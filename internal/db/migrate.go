package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/santzin/gin-tattoo/internal/data"
)

var validSchema = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)

const ddl = `
CREATE TABLE IF NOT EXISTS styles (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    origin      TEXT NOT NULL,
    popularity  TEXT NOT NULL CHECK (popularity IN ('high', 'medium', 'low'))
);

CREATE TABLE IF NOT EXISTS curiosities (
    id       SERIAL PRIMARY KEY,
    title    TEXT NOT NULL,
    content  TEXT NOT NULL,
    category TEXT NOT NULL CHECK (category IN ('history', 'culture', 'science', 'art'))
);
`

// Migrate ensures the schema exists, applies DDL, and seeds initial data.
func Migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if schema := os.Getenv("DB_SCHEMA"); schema != "" {
		if !validSchema.MatchString(schema) {
			return fmt.Errorf("invalid DB_SCHEMA value: %q", schema)
		}
		if _, err := pool.Exec(ctx, fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)); err != nil {
			return fmt.Errorf("failed to create schema %q: %w", schema, err)
		}
	}

	if _, err := pool.Exec(ctx, ddl); err != nil {
		return fmt.Errorf("failed to apply DDL: %w", err)
	}

	if err := seedStyles(ctx, pool); err != nil {
		return err
	}
	return seedCuriosities(ctx, pool)
}

func seedStyles(ctx context.Context, pool *pgxpool.Pool) error {
	var count int
	if err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM styles").Scan(&count); err != nil {
		return fmt.Errorf("failed to count styles: %w", err)
	}
	if count > 0 {
		return nil
	}
	log.Println("db: seeding styles")
	for _, s := range data.Styles {
		if _, err := pool.Exec(ctx,
			"INSERT INTO styles (name, description, origin, popularity) VALUES ($1, $2, $3, $4)",
			s.Name, s.Description, s.Origin, s.Popularity,
		); err != nil {
			return fmt.Errorf("seed style %q: %w", s.Name, err)
		}
	}
	return nil
}

func seedCuriosities(ctx context.Context, pool *pgxpool.Pool) error {
	var count int
	if err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM curiosities").Scan(&count); err != nil {
		return fmt.Errorf("failed to count curiosities: %w", err)
	}
	if count > 0 {
		return nil
	}
	log.Println("db: seeding curiosities")
	for _, c := range data.Curiosities {
		if _, err := pool.Exec(ctx,
			"INSERT INTO curiosities (title, content, category) VALUES ($1, $2, $3)",
			c.Title, c.Content, c.Category,
		); err != nil {
			return fmt.Errorf("seed curiosity %q: %w", c.Title, err)
		}
	}
	return nil
}
