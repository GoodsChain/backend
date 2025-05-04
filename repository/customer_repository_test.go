package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GoodsChain/backend/model"
	"github.com/jmoiron/sqlx"
)

// Helper function to create a new mock database for testing
func newMockDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return sqlxDB, mock
}

func TestNewCustomerRepository(t *testing.T) {
	db, _ := newMockDB(t)
	repo := NewCustomerRepository(db)

	if repo == nil {
		t.Fatal("Expected repository to be created, got nil")
	}
}

func TestCreate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerRepository(db)

	customer := &model.Customer{
		ID:      "cust123",
		Name:    "Test Customer",
		Address: "123 Test St",
		Phone:   "+1234567890",
		Email:   "test@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO customer \\(id, name, address, phone, email, created_by, updated_by\\)").
			WithArgs(customer.ID, customer.Name, customer.Address, customer.Phone, customer.Email, "admin", "admin").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(customer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("INSERT INTO customer").
			WithArgs(customer.ID, customer.Name, customer.Address, customer.Phone, customer.Email, "admin", "admin").
			WillReturnError(expectedErr)

		err := repo.Create(customer)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestGet(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerRepository(db)

	customerID := "cust123"
	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow(customerID, "Test Customer", "123 Test St", "+1234567890", "test@example.com", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by FROM customer WHERE id = \\$1").
			WithArgs(customerID).
			WillReturnRows(rows)

		customer, err := repo.Get(customerID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if customer == nil {
			t.Fatal("Expected customer, got nil")
		}

		if customer.ID != customerID {
			t.Errorf("Expected ID %s, got %s", customerID, customer.ID)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		expectedErr := errors.New("no rows in result set")
		mock.ExpectQuery("SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by FROM customer WHERE id = \\$1").
			WithArgs(customerID).
			WillReturnError(expectedErr)

		customer, err := repo.Get(customerID)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if customer != nil && customer.ID != "" {
			t.Errorf("Expected nil or empty customer, got %v", customer)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerRepository(db)

	customerID := "cust123"
	customer := &model.Customer{
		ID:      customerID,
		Name:    "Updated Customer",
		Address: "456 New St",
		Phone:   "+9876543210",
		Email:   "updated@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("UPDATE customer SET name = \\$1, address = \\$2, phone = \\$3, email = \\$4, updated_by = \\$5, updated_at = now\\(\\) WHERE id = \\$6").
			WithArgs(customer.Name, customer.Address, customer.Phone, customer.Email, "admin", customerID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(customerID, customer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("UPDATE customer SET").
			WithArgs(customer.Name, customer.Address, customer.Phone, customer.Email, "admin", customerID).
			WillReturnError(expectedErr)

		err := repo.Update(customerID, customer)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestDelete(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerRepository(db)

	customerID := "cust123"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM customer WHERE id = \\$1").
			WithArgs(customerID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Delete(customerID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("DELETE FROM customer WHERE id = \\$1").
			WithArgs(customerID).
			WillReturnError(expectedErr)

		err := repo.Delete(customerID)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestGetAll(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerRepository(db)

	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success With Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow("cust123", "Customer 1", "123 Test St", "+1234567890", "cust1@example.com", createdAt, "admin", updatedAt, "admin").
			AddRow("cust456", "Customer 2", "456 Test St", "+0987654321", "cust2@example.com", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT \\* FROM customer").
			WillReturnRows(rows)

		customers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(customers) != 2 {
			t.Errorf("Expected 2 customers, got %d", len(customers))
		}

		if customers[0].ID != "cust123" || customers[1].ID != "cust456" {
			t.Errorf("Unexpected customer IDs: %v, %v", customers[0].ID, customers[1].ID)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Success With No Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"})

		mock.ExpectQuery("SELECT \\* FROM customer").
			WillReturnRows(rows)

		customers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// When no rows are found, the library might return nil or an empty slice
		// Both are acceptable as long as we can safely iterate over the result
		if customers == nil || len(customers) == 0 {
			// This is fine, either nil or empty slice is acceptable
		} else {
			t.Error("Expected nil or empty slice")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectQuery("SELECT \\* FROM customer").
			WillReturnError(expectedErr)

		customers, err := repo.GetAll()
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if customers != nil {
			t.Errorf("Expected nil result, got %v", customers)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}
