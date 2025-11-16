package storage

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
)

func (pgPepo *PostgresRepository) GetUserByID(ctx context.Context, userID domain.UserId) (*domain.User, error) {
	const selectUserByIDQ = `
SELECT u.id,
	u.username,
	u.team_id,
	t.name AS team_name,
	u.is_active
FROM users u
JOIN teams t ON u.team_id = t.id
WHERE u.id = $1
    `
	//достаём юзера
	var tmpUser domain.User
	err := pgPepo.db.QueryRowContext(ctx, selectUserByIDQ, userID).Scan(
		&tmpUser.UserId,
		&tmpUser.Username,
		&tmpUser.TeamName,
		&tmpUser.IsActive,
	)
	if err != nil {
		if IsErrorPGNotExist(err) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return domain.NewUser(
		tmpUser.UserId, tmpUser.Username,
		tmpUser.TeamName, tmpUser.IsActive,
	)
}
func (pgPepo *PostgresRepository) UpdateUser(ctx context.Context, user domain.User) error {
	const selectTeamIDByTeamNameQ = `
SELECT id
FROM teams
WHERE name = $1
`
	var teamId int
	err := pgPepo.db.QueryRowContext(ctx, selectTeamIDByTeamNameQ, user).Scan(
		&teamId,
	)
	if err != nil {
		if IsErrorPGNotExist(err) {
			return domain.ErrTeamNotFound
		}
		return err
	}

	const updateUserQ = `
UPDATE users
SET
    username = $2,
    team_id  = $3,
    is_active = $4
WHERE id = $1
`

	res, err := pgPepo.db.ExecContext(ctx, updateUserQ,
		user.UserId,
		user.Username,
		teamId,
		user.IsActive,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
func (pgPepo *PostgresRepository) CreateAndUpdateUsers(ctx context.Context, users []domain.User) error {
	if len(users) == 0 {
		return nil
	}

	tx, err := pgPepo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const insertOrUpdateUserQ = `
INSERT INTO users (id, username, team_id, is_active)
SELECT $1, $2, t.id, $4
FROM teams t
WHERE t.name = $3
ON CONFLICT (id) DO UPDATE SET
    username  = EXCLUDED.username,
    team_id   = EXCLUDED.team_id,
    is_active = EXCLUDED.is_active
`

	for _, u := range users {
		res, err := tx.ExecContext(ctx, insertOrUpdateUserQ,
			u.UserId,
			u.Username,
			u.TeamName,
			u.IsActive,
		)
		if err != nil {
			return err
		}

		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			return domain.ErrTeamNotFound
		}
	}

	return tx.Commit()
}
func (pgPepo *PostgresRepository) ListUsersByTeamName(ctx context.Context, teamName string) ([]domain.User, error) {
	const selectTeamIDByTeamNameQ = `
SELECT id
FROM teams
WHERE name = $1
`
	var teamID int
	err := pgPepo.db.QueryRowContext(ctx, selectTeamIDByTeamNameQ, teamName).Scan(&teamID)
	if err != nil {
		if IsErrorPGNotExist(err) {
			return nil, domain.ErrTeamNotFound
		}
		return nil, err
	}

	//Достаём всех пользователей по teamID
	const selectUsersByTeamIdQ = `
SELECT id, username, is_active
FROM users
WHERE team_id = $1
`
	rows, err := pgPepo.db.QueryContext(ctx, selectUsersByTeamIdQ, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)

	for rows.Next() {
		var tmpUser domain.User
		if err := rows.Scan(&tmpUser.UserId, &tmpUser.Username, &tmpUser.IsActive); err != nil {
			return nil, err
		}
		//если юзер битый, пропускаем
		newUser, err := domain.NewUser(tmpUser.UserId, tmpUser.Username, teamName, tmpUser.IsActive)
		if err != nil {
			continue
		}
		users = append(users, *newUser)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
