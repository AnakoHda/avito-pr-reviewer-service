package teamHandler

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	AddTeamWithUsers(ctx context.Context, team domain.Team) error
	getTeam(ctx context.Context, teamName string) (*domain.Team, error)
}
type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) AddTeam(w http.ResponseWriter, r *http.Request) {
	var req TeamDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse := ErrorResponse{
			ErrorBody{
				Code:    ErrTeamExists,
				Message: "team_name already exist",
			},
		}
		_ = json.NewEncoder(w).Encode(jsonResponse)
		return
	}
	
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) GetTeam(w http.ResponseWriter, r *http.Request) {

}

func FromTeamDTOToTeam(dto TeamDTO) (domain.Team, error) {
	var tmpTeam domain.Team
	tmpTeam.TeamName = dto.TeamName
	for _, member := range dto.Members {
		tmpUser, err := domain.NewUser(domain.UserId(member.UserID), member.Username, dto.TeamName, member.IsActive)
		if err != nil {
			return domain.Team{}, err
		}
		tmpTeam.Members = append(tmpTeam.Members, *tmpUser)
	}
	return tmpTeam, nil
}
