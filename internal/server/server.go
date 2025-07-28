package server

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	profile *Profile
	server  *http.Server
}

func NewServer(profile *Profile) (s *Server, err error) {
	s = &Server{
		profile: profile,
	}
	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	handler := registerHandlers(s.profile.Container)
	s.server = &http.Server{
		Addr:    s.profile.Addr,
		Handler: h2c.NewHandler(handler, &http2.Server{}),
	}
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()
	return nil
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Info("shutdown...")
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Errorf("failed to shutdown server, error: %v\n", err)
	}
}
