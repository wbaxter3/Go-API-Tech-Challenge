package main

import (
	"context"
	"errors"
	"fmt"
	"go-api-tech-challenge/internal/config"
	"go-api-tech-challenge/internal/routes"
	"go-api-tech-challenge/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Startup failed. err: %v", err)
	}
}

func run() error {
	// Setup
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("[in run]: %w", err)
	}

	logger := httplog.NewLogger("api", httplog.Options{
		LogLevel:        cfg.LogLevel,
		JSON:            false,
		Concise:         true,
		ResponseHeaders: false,
	})

	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(logger))
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		MaxAge:         300,
	}))

	svs := services.NewCourseService(db)
	routes.RegisterRoutes(router, logger, svs, routes.WithRegisterHealthRoute(true))

	//if cfg.HTTPUseSwagger {
	//swagger.RunSwagger(router, logger, cfg.HTTPDomain+cfg.HTTPPort)
	//}

	serverInstance := &http.Server{
		Addr:              cfg.HTTPDomain + cfg.HTTPPort,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		WriteTimeout:      500 * time.Millisecond,
		Handler:           router,
	}

	// Graceful shutdown
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		fmt.Println()
		logger.Info("Shutdown signal received")

		shutdownCtx, err := context.WithTimeout(
			serverCtx, time.Duration(cfg.HTTPShutdownDuration)*time.Second,
		)
		if err != nil {
			log.Fatalf("Error creating context.WithTimeout. err: %v", err)
		}

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		if err := serverInstance.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Error shutting down server. err: %v", err)
		}
		serverStopCtx()
	}()

	// Run
	logger.Info(fmt.Sprintf("Server is listening on %s", serverInstance.Addr))
	err = serverInstance.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	<-serverCtx.Done()
	logger.Info("Shutdown complete")
	return nil
}
