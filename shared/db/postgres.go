package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	Conn *pgx.Conn
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(ctx context.Context, dbURL string) (*PostgresDB, error) {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{Conn: conn}, nil
}

// Connect is a helper function for direct connection
func Connect(ctx context.Context, dbURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Close closes the database connection
func (db *PostgresDB) Close(ctx context.Context) error {
	return db.Conn.Close(ctx)
}
