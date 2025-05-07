package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/GoodsChain/backend/model"
	"github.com/jmoiron/sqlx"
)

// CustomerCarRepository defines the interface for customer_car data operations
type CustomerCarRepository interface {
	Create(customerCar *model.CustomerCar) error
	GetByID(id string) (*model.CustomerCar, error)
	GetAll() ([]*model.CustomerCar, error)
	GetByCustomerID(customerID string) ([]*model.CustomerCar, error)
	GetByCarID(carID string) ([]*model.CustomerCar, error)
	Update(id string, customerCar *model.CustomerCar) error
	Delete(id string) error
}

type customerCarRepository struct {
	db *sqlx.DB
}

// NewCustomerCarRepository creates a new instance of CustomerCarRepository
func NewCustomerCarRepository(db *sqlx.DB) CustomerCarRepository {
	return &customerCarRepository{db: db}
}

// Create adds a new customer_car relationship to the database
func (r *customerCarRepository) Create(customerCar *model.CustomerCar) error {
	customerCar.CreatedAt = time.Now()
	customerCar.UpdatedAt = time.Now()
	// CreatedBy and UpdatedBy should be set by the application/usecase layer

	query := `INSERT INTO customer_car (id, car_id, cust_id, created_at, created_by, updated_at, updated_by)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, customerCar.ID, customerCar.CarID, customerCar.CustomerID, 
		customerCar.CreatedAt, customerCar.CreatedBy, customerCar.UpdatedAt, customerCar.UpdatedBy)
	return err
}

// GetByID retrieves a customer_car relationship by its ID
func (r *customerCarRepository) GetByID(id string) (*model.CustomerCar, error) {
	var customerCar model.CustomerCar
	query := `SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by 
	          FROM customer_car WHERE id = $1`
	err := r.db.Get(&customerCar, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &customerCar, nil
}

// GetAll retrieves all customer_car relationships from the database
func (r *customerCarRepository) GetAll() ([]*model.CustomerCar, error) {
	var customerCars []*model.CustomerCar
	query := `SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by 
	          FROM customer_car ORDER BY created_at DESC`
	err := r.db.Select(&customerCars, query)
	if err != nil {
		return nil, err
	}
	return customerCars, nil
}

// GetByCustomerID retrieves all customer_car relationships for a specific customer
func (r *customerCarRepository) GetByCustomerID(customerID string) ([]*model.CustomerCar, error) {
	var customerCars []*model.CustomerCar
	query := `SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by 
	          FROM customer_car WHERE cust_id = $1 ORDER BY created_at DESC`
	err := r.db.Select(&customerCars, query, customerID)
	if err != nil {
		return nil, err
	}
	return customerCars, nil
}

// GetByCarID retrieves all customer_car relationships for a specific car
func (r *customerCarRepository) GetByCarID(carID string) ([]*model.CustomerCar, error) {
	var customerCars []*model.CustomerCar
	query := `SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by 
	          FROM customer_car WHERE car_id = $1 ORDER BY created_at DESC`
	err := r.db.Select(&customerCars, query, carID)
	if err != nil {
		return nil, err
	}
	return customerCars, nil
}

// Update updates an existing customer_car relationship
func (r *customerCarRepository) Update(id string, customerCar *model.CustomerCar) error {
	customerCar.UpdatedAt = time.Now()
	// UpdatedBy should be set by the application/usecase layer

	query := `UPDATE customer_car SET car_id = $1, cust_id = $2, updated_at = $3, updated_by = $4 WHERE id = $5`
	result, err := r.db.Exec(query, customerCar.CarID, customerCar.CustomerID, 
		customerCar.UpdatedAt, customerCar.UpdatedBy, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// Delete removes a customer_car relationship from the database by its ID
func (r *customerCarRepository) Delete(id string) error {
	query := `DELETE FROM customer_car WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
