package storage

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func IsErrorPGAlreadyExist(err error) bool {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}
func IsErrorPGNotExist(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	return false
}
