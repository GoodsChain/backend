package main

import (
	"fmt"
	"github.com/GoodsChain/backend/config"
	"github.com/GoodsChain/backend/handler"
	"github.com/GoodsChain/backend/repository"
	"github.com/GoodsChain/backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	customerRepo := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	supplierRepo := repository.NewSupplierRepository(db)
	supplierUsecase := usecase.NewSupplierUsecase(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierUsecase)

	handler.InitRoutes(r, customerHandler, supplierHandler)

	log.Printf("Server starting on port %s", cfg.APIPort)
	if err := r.Run(":" + cfg.APIPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
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
