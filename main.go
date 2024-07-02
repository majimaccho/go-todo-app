package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/majimaccho/go-todo-app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("Failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen port %v", err)
	}
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
