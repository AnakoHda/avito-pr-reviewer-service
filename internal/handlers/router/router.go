package router

import (
	"avito-pr-reviewer-service/internal/handlers/pullRequestHandler"
	"avito-pr-reviewer-service/internal/handlers/teamHandler"
	"avito-pr-reviewer-service/internal/handlers/userHandler"
	"avito-pr-reviewer-service/internal/service/pullRequestService"
	"avito-pr-reviewer-service/internal/service/teamService"
	"avito-pr-reviewer-service/internal/service/userService"
	"net/http"
)

func RegisterRoutes(prServ *pullRequestService.Service, teamServ *teamService.Service, userServ *userService.Service) http.Handler {
	mux := http.NewServeMux()
	pullRequestHandler.New(prServ).Register(mux)
	teamHandler.New(teamServ).Register(mux)
	userHandler.New(userServ).Register(mux)
	return mux
}
