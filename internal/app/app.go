package app

import (
	"avito-pr-reviewer-service/internal/handlers/router"
	"avito-pr-reviewer-service/internal/service/pullRequestService"
	"avito-pr-reviewer-service/internal/service/teamService"
	"avito-pr-reviewer-service/internal/service/userService"
	"avito-pr-reviewer-service/internal/storage"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := storage.NewConnection(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("POSTGRES Connected")

	defer db.Close()

	serverCtx, stopServer := context.WithCancel(context.Background())
	defer stopServer()

	pero := storage.New(db)
	prServ := pullRequestService.New(pero, pero)
	teamServ := teamService.New(pero, pero)
	userServ := userService.New(pero, pero)

	ServerPORT := os.Getenv("SERVICE_PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServerPORT),
		Handler: router.RegisterRoutes(prServ, teamServ, userServ),
		BaseContext: func(net.Listener) context.Context {
			return serverCtx
		},
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		return err
	}
	fmt.Println("Termination")
	return nil
}
