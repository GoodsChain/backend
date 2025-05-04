package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/GoodsChain/backend/config"
	"github.com/GoodsChain/backend/handler"
	"github.com/GoodsChain/backend/repository"
	"github.com/GoodsChain/backend/usecase"
)

func main() {
	// Initialize application context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	cfg := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Connect to database
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories, usecases, and handlers
	customerRepo := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	supplierRepo := repository.NewSupplierRepository(db)
	supplierUsecase := usecase.NewSupplierUsecase(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierUsecase)

	// Initialize routes
	handler.InitRoutes(r, customerHandler, supplierHandler)

	// Configure HTTP server with reasonable timeouts
	srv := &http.Server{
		Addr:         ":" + cfg.APIPort,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so it doesn't block signal handling
	go func() {
		log.Printf("Server starting on port %s", cfg.APIPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Start the graceful shutdown handling
	handleGracefulShutdown(ctx, srv, db)
}

// handleGracefulShutdown manages the graceful shutdown process for the server
func handleGracefulShutdown(ctx context.Context, srv *http.Server, db *sqlx.DB) {
	// Set up channel to listen for signals
	quit := make(chan os.Signal, 1)
	// Listen for SIGINT and SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-quit
	log.Printf("Received signal: %v. Shutting down server...", sig)

	// Create a deadline for server shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	log.Println("Closing database connection...")
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
	log.Println("Database connection closed")

	log.Println("Server exited gracefully")
}

func connectDB(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=10",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
