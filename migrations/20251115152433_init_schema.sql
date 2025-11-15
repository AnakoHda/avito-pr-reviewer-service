-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE users (
    id           TEXT PRIMARY KEY,
    username     TEXT NOT NULL,
    team_id      INT NOT NULL REFERENCES teams(id) ON DELETE RESTRICT,
    is_active    BOOLEAN NOT NULL
);

CREATE TABLE teams (
    id      SERIAL PRIMARY KEY,
    name    TEXT NOT NULL UNIQUE
);

CREATE TYPE pull_request_status AS ENUM ('OPEN', 'MERGED');
CREATE TABLE pull_requests (
    id                  TEXT PRIMARY KEY,
    pull_requests_name  TEXT NOT NULL,
    author_id           TEXT NOT NULL REFERENCES users(user_id),
    status              pull_request_status NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    merged_at           TIMESTAMPTZ NULL
);

CREATE TABLE pull_request_reviewers (
    pull_request_id     TEXT NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    reviewer_id         TEXT NOT NULL REFERENCES users(user_id),
    PRIMARY KEY (pull_request_id, reviewer_id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

DROP TYPE IF EXISTS pull_request_status;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS pull_request_reviewers;