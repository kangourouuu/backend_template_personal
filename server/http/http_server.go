package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	serviceName = "backend_template"
	version     = "0.0.1"
)

// Server wraps the http.Server to manage the server's lifecycle.
type Server struct {
	httpServer *http.Server
}

// New creates a new server instance. It takes a port and a Gin engine.
func New(port string, engine *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: engine,
		},
	}
}

// NewEngine initializes a new gin.Engine with default middleware and a health check endpoint.
func NewEngine() *gin.Engine {
	engine := gin.New()
	engine.GET("/service", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"service": serviceName,
			"version": version,
		})
	})
	return engine
}

// Run starts the server and blocks until a shutdown signal is received.
// It handles the graceful shutdown of the server.
func (s *Server) Run() error {
	go func() {
		logrus.Infof("Service '%s' listening on %s", serviceName, s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	// The context is used to inform the server it has 30 seconds to finish
	// the requests it is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logrus.Errorf("Server forced to shutdown: %v", err)
		return err
	}

	logrus.Info("Server exiting")
	return nil
}
