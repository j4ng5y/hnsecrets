package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Option is a function type that defines what functional options are available to build a new Server instance
type Option func(*Server)

func WithIPAddress(ip string) Option {
	return func(S *Server) {
		S.address = ip
	}
}

func WithPort(port int) Option {
	return func(S *Server) {
		S.port = port
	}
}

func WithDefaults() Option {
	return func(S *Server) {
		S.HTTP.Addr = "0.0.0.0:8080"
		S.HTTP.IdleTimeout = 60 * time.Second
		S.HTTP.ReadHeaderTimeout = 15 * time.Second
		S.HTTP.ReadTimeout = 15 * time.Second
		S.HTTP.WriteTimeout = 15 * time.Second
	}
}

// Server is the primary struct that holds and handles all data related to core operation.
type Server struct {
	address string
	port    int
	HTTP    *http.Server
}

func (S *Server) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	runChan := make(chan os.Signal, 1)

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	log.Printf("Running server on %s\n", S.HTTP.Addr)
	go func() {
		S.HTTP.ListenAndServe()
	}()

	sig := <-runChan

	log.Printf("Stopping server due to %s\n", sig.String())
	err := S.HTTP.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func New(opts ...Option) (S *Server, err error) {
	S = &Server{
		HTTP: &http.Server{},
	}
	for _, opt := range opts {
		opt(S)
	}
	S.HTTP.Handler = newRouter()
	return S, err
}
