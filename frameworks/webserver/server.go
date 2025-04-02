package webserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/charlitoro/go-clean-architecture-skeleton/frameworks/webserver/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router           *gin.Engine
	port             string
	statusController *controllers.StatusController
}

// NewServer creates a new server instance
func NewServer(port string, statusController *controllers.StatusController) *Server {
	router := gin.Default()

	return &Server{
		router:           router,
		port:             port,
		statusController: statusController,
	}
}

// SetupRoutes registers all the routes for the application
func (s *Server) SetupRoutes() {
	// Setup status routes
	routes.SetupStatusRoutes(s.router, s.statusController)

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
		log.Println("Shutting down server...")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
		wg.Done()
	}()

	// Start the server
	log.Printf("Server is running on port %s\n", s.port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for server context to be stopped
	wg.Wait()
	log.Println("Server stopped")
}
