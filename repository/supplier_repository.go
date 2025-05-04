package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/GoodsChain/backend/model"
)

type SupplierRepository interface {
	Create(supplier *model.Supplier) error
	Get(id string) (*model.Supplier, error)
	Update(id string, supplier *model.Supplier) error
	Delete(id string) error
	GetAll() ([]*model.Supplier, error)
}

type supplierRepository struct {
	db *sqlx.DB
}

func NewSupplierRepository(db *sqlx.DB) SupplierRepository {
	return &supplierRepository{
		db: db,
	}
}

func (r *supplierRepository) Create(supplier *model.Supplier) error {
	query := `INSERT INTO supplier (id, name, address, phone, email, created_by, updated_by) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, supplier.ID, supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", "admin")
	return err
}

func (r *supplierRepository) Get(id string) (*model.Supplier, error) {
	var supplier model.Supplier
	query := `SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by 
		FROM supplier WHERE id = $1`
	err := r.db.Get(&supplier, query, id)
	return &supplier, err
}

func (r *supplierRepository) Update(id string, supplier *model.Supplier) error {
	query := `UPDATE supplier SET name = $1, address = $2, phone = $3, email = $4, updated_by = $5, updated_at = now() 
		WHERE id = $6`
	_, err := r.db.Exec(query, supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", id)
	return err
}

func (r *supplierRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM supplier WHERE id = $1", id)
	return err
}

func (r *supplierRepository) GetAll() ([]*model.Supplier, error) {
	var suppliers []*model.Supplier
	if err := r.db.Select(&suppliers, "SELECT * FROM supplier"); err != nil {
		return nil, err
	}
	return suppliers, nil
}
