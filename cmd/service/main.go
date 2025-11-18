package main

import (
	"avito-pr-reviewer-service/internal/app"
	"context"
	"log/slog"
	"os"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		slog.Error("app exited with error: %v", err)
		os.Exit(1)
	}
}
