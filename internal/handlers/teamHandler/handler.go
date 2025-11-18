package teamHandler

import (
	"avito-pr-reviewer-service/internal/domain"
	"avito-pr-reviewer-service/internal/generated/api/dto"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Service interface {
	AddTeamWithUsers(ctx context.Context, team domain.Team) (*domain.Team, error)
	GetTeam(ctx context.Context, teamName string) (*domain.Team, error)
}
type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/team/add", h.POSTAddTeam)
	mux.HandleFunc("/team/get", h.GETGetTeam)
}

func (h *Handler) POSTAddTeam(w http.ResponseWriter, r *http.Request) {
	var req dto.Team
	slog.Info("Touch ADD TEAM")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseFormatError(w, http.StatusBadRequest, ErrBadRequest, "Decode error")
		return
	}
	tmpTeam, err := FromTeamDTOToTeam(req)
	if err != nil {
		ResponseFormatError(w, http.StatusBadRequest, ErrBadRequest, "Format error")
		return
	}
	ctx := r.Context()
	actualTeam, err := h.svc.AddTeamWithUsers(ctx, tmpTeam)
	if err != nil {
		if err.Error() == domain.ErrTeamAlreadyExists.Error() {
			ResponseFormatError(w, http.StatusBadRequest, dto.TEAMEXISTS, "team_name already exists")
			return
		}
		ResponseFormatError(w, http.StatusInternalServerError, ErrInternal, "service error")
		return
	}
	responseTeamDTO := FromTeamToDTO(*actualTeam)
	ResponseFormatOK(w, http.StatusCreated, responseTeamDTO)
	return
}

func (h *Handler) GETGetTeam(w http.ResponseWriter, r *http.Request) {

}
