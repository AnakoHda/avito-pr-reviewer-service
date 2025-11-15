package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Termination")
	return nil
}
