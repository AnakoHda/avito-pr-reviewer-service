package storage

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
)

func (pgPepo *PostgresRepository) CreateTeam(ctx context.Context, teamName string) error {
	const insertTeamQ = `
		INSERT INTO teams (name)
		VALUES ($1)
	`

	_, err := pgPepo.db.ExecContext(ctx, insertTeamQ, teamName)
	if err != nil {
		//если существует такой TeamName
		if IsErrorPGAlreadyExist(err) {
			return domain.ErrTeamAlreadyExists
		}
		return err
	}

	return nil
}
