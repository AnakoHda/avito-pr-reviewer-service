package teamHandler

type TeamMemberDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type TeamDTO struct {
	TeamName string          `json:"team_name"`
	Members  []TeamMemberDTO `json:"members"`
}

const (
	ErrTeamExists  = "TEAM_EXISTS"
	ErrPRExists    = "PR_EXISTS"
	ErrPRMerged    = "PR_MERGED"
	ErrNotAssigned = "NOT_ASSIGNED"
	ErrNoCandidate = "NO_CANDIDATE"
	ErrNotFound    = "NOT_FOUND"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
