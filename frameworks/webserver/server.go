package webserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/services"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router           *gin.Engine
	port             string
	logger           *services.Logger
}

// NewServer creates a new server instance
func NewServer(port string, logger *services.Logger) *Server {
	router := gin.Default()

	return &Server{
		router:           router,
		port:             port,
		logger:           logger,
	}
}

// SetupRoutes registers all the routes for the application
func (s *Server) SetupRoutes() {
	// Setup status routes
	routes.StatusRoutes(s.router)
	routes.PostRoutes(s.router)

	// TODO: Register additional routes here
}

// Start starts the HTTP server
func (s *Server) Start() {
	// Setup routes
	s.SetupRoutes()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.port),
		Handler: s.router,
	}

	// Use WaitGroup to prevent the function from returning until the server is shut down
	var wg sync.WaitGroup
	wg.Add(1)

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		s.logger.Info("Shutting down server...")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				s.logger.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		if err := server.Shutdown(shutdownCtx); err != nil {
			s.logger.Fatal("Shutdown error", zap.Error(err))
		}
		serverStopCtx()
		wg.Done()
	}()

	// Start the server
	s.logger.Info("Server is running", zap.String("port", s.port))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Server error", zap.Error(err))
		}
	}()

	// Wait for server context to be stopped
	wg.Wait()
	s.logger.Info("Server stopped")
}
