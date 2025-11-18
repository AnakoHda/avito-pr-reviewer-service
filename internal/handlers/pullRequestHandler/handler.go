package pullRequestHandler

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
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

}

func (h *Handler) POSTMergePullRequest(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) POSTReassignPullRequest(w http.ResponseWriter, r *http.Request) {}
