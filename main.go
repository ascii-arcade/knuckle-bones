package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ascii-arcade/game-template/app"
	"github.com/ascii-arcade/game-template/config"
	"github.com/ascii-arcade/game-template/web"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(config.Host, config.SSHPort)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(app.TeaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("could not create wish server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("starting ssh server", "host", config.Host, "port", config.SSHPort, "version", config.Version)
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("could not start wish server", "error", err)
			done <- nil
		}
	}()

	go func() {
		log.Info("starting http server", "host", config.Host, "port", config.HTTPPort, "version", config.Version)
		if err := web.Run(); err != nil {
			log.Error("could not start web server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("shutting down servers...")
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("error shutting down servers", "error", err)
	}
}
