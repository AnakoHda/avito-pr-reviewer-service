package userHandler

import (
	"avito-pr-reviewer-service/internal/domain"
	"context"
	"net/http"
)

type Service interface {
	SetIsActive(ctx context.Context, userId domain.UserId, isActive bool) (*domain.User, error)
	GetReview(ctx context.Context, userId domain.UserId) ([]domain.PullRequest, error)
}
type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/users/serIsActive", h.POSTSetIsActiveUser)
	mux.HandleFunc("/users/getReview", h.GETGetReviewUser)
}

func (h *Handler) POSTSetIsActiveUser(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GETGetReviewUser(w http.ResponseWriter, r *http.Request) {}
