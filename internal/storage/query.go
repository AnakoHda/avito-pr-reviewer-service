package storage

const selectPullRequestIDsQ = `
SELECT pull_request_id
FROM pull_request_reviewers
WHERE reviewer_id = $1
`
const pullRequestSelectQ = `
SELECT id, pull_requests_name, author_id, status, created_at, merged_at
FROM pull_requests
WHERE id = $1
`
const reviewersSelectQ = `
SELECT reviewer_id
FROM pull_request_reviewers
WHERE pull_request_id = $1
`
const pullRequestInsertQ = `
INSERT INTO pull_requests (id, pull_requests_name, author_id, status, created_at, merged_at)
VALUES ($1, $2, $3, $4, $5, $6)
`
const pullRequestUpdateQ = `
UPDATE pull_requests
SET
    pull_requests_name = $2,
    author_id          = $3,
    status             = $4,
    created_at         = $5,
    merged_at          = $6
WHERE id = $1
`
const deleteReviewersQ = `
DELETE FROM pull_request_reviewers
WHERE pull_request_id = $1
`
const insertReviewerQ = `
INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)
VALUES ($1, $2)
`

// team

const insertTeamQ = `
		INSERT INTO teams (name)
		VALUES ($1)
		RETURNING id
	`
const insertOrUpdateUserQ = `
	INSERT INTO users (id, username, team_id, is_active)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (id) DO UPDATE SET
		username  = EXCLUDED.username,
		team_id   = EXCLUDED.team_id,
		is_active = EXCLUDED.is_active
	`

// users

const selectUserByIDQ = `
SELECT u.id,
	u.username,
	t.name AS team_name,
	u.is_active
FROM users u
JOIN teams t ON u.team_id = t.id
WHERE u.id = $1
    `
const updateUserQ = `
UPDATE users
SET
    username = $2,
    team_id  = $3,
    is_active = $4
WHERE id = $1
`
const selectTeamIDByTeamNameQ = `
SELECT id
FROM teams
WHERE name = $1
`
const selectUsersByTeamIdQ = `
SELECT id, username, is_active
FROM users
WHERE team_id = $1
`
