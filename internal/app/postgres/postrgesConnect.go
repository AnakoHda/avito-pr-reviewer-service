package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InitPostgres(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
