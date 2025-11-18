package pullRequestHandler

import (
	"avito-pr-reviewer-service/internal/domain"
	"avito-pr-reviewer-service/internal/generated/api/dto"
	"avito-pr-reviewer-service/internal/handlers"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Service interface {
	CreatePullRequest(ctx context.Context, PullRequestId domain.PullRequestId, PullRequestName string, AuthorId domain.UserId) (*domain.PullRequest, error)
	Merge(ctx context.Context, pullRequestID domain.PullRequestId) (*domain.PullRequest, error)
	Reassign(ctx context.Context, pullRequestID domain.PullRequestId, oldReviewerId domain.UserId) (*domain.PullRequest, error)
}
type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/pullRequest/create", h.POSTCreatePullRequest)
	mux.HandleFunc("/pullRequest/merge", h.POSTMergePullRequest)
	mux.HandleFunc("/pullRequest/reassign", h.POSTReassignPullRequest)
}

func (h *Handler) POSTCreatePullRequest(w http.ResponseWriter, r *http.Request) {
	var req dto.PostPullRequestCreateJSONBody
	slog.Info("Touch CREATE PULL REQUEST")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handlers.ResponseFormatError(w, http.StatusBadRequest, handlers.ErrBadRequest, "Decode error")
		return
	}

	ctx := r.Context()
	createdPR, err := h.svc.CreatePullRequest(
		ctx,
		domain.PullRequestId(req.PullRequestId),
		req.PullRequestName, domain.UserId(req.AuthorId),
	)
	if err != nil {
		if err.Error() == domain.ErrUserNotFound.Error() || err.Error() == domain.ErrTeamNotFound.Error() {
			handlers.ResponseFormatError(w, http.StatusNotFound, dto.NOTFOUND, "resource not found")
			return
		}
		if err.Error() == domain.ErrPullRequestAlreadyExists.Error() {
			handlers.ResponseFormatError(w, http.StatusConflict, dto.PREXISTS, "PR id already exists")
			return
		}
		handlers.ResponseFormatError(w, http.StatusInternalServerError, handlers.ErrInternal, "service error")
		return
	}
	assigned := make([]string, len(createdPR.AssignedReviewers))
	for i, reviewer := range createdPR.AssignedReviewers {
		assigned[i] = string(reviewer)
	}
	handlers.ResponseFormatOK(
		w,
		http.StatusCreated,
		handlers.PullRequestCreateResponse{
			PR: handlers.PullRequestCreate{
				PullRequestId:     string(createdPR.PullRequestId),
				PullRequestName:   createdPR.PullRequestName,
				AuthorId:          string(createdPR.AuthorId),
				Status:            dto.PullRequestStatus(createdPR.Status),
				AssignedReviewers: assigned,
			}})
	return
}

func (h *Handler) POSTMergePullRequest(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) POSTReassignPullRequest(w http.ResponseWriter, r *http.Request) {}
