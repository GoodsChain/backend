package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/GoodsChain/backend/model"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	Get(id string) (*model.Customer, error)
	Update(id string, customer *model.Customer) error
	Delete(id string) error
	GetAll() ([]*model.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (r *customerRepository) GetAll() ([]*model.Customer, error) {
	var customers []*model.Customer
	if err := r.db.Select(&customers, "SELECT * FROM customer"); err != nil { // Using singular form to match table name
		return nil, err
	}
	return customers, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Create(customer *model.Customer) error {
	query := `INSERT INTO customer (id, name, address, phone, email, created_by, updated_by) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, customer.ID, customer.Name, customer.Address, customer.Phone, customer.Email, "admin", "admin")
	return err
}

func (r *customerRepository) Get(id string) (*model.Customer, error) {
	var customer model.Customer
	query := `SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by 
		FROM customer WHERE id = $1`
	err := r.db.Get(&customer, query, id)
	return &customer, err
}

func (r *customerRepository) Update(id string, customer *model.Customer) error {
	query := `UPDATE customer SET name = $1, address = $2, phone = $3, email = $4, updated_by = $5, updated_at = now() 
		WHERE id = $6`
	_, err := r.db.Exec(query, customer.Name, customer.Address, customer.Phone, customer.Email, "admin", id)
	return err
}

func (r *customerRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM customer WHERE id = $1", id)
	return err
}
