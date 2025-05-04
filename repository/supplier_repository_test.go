package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GoodsChain/backend/model"
)

func TestNewSupplierRepository(t *testing.T) {
	db, _ := newMockDB(t)
	repo := NewSupplierRepository(db)

	if repo == nil {
		t.Fatal("Expected repository to be created, got nil")
	}
}

func TestSupplierCreate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSupplierRepository(db)

	supplier := &model.Supplier{
		ID:      "supp123",
		Name:    "Test Supplier",
		Address: "123 Test St",
		Phone:   "+1234567890",
		Email:   "supplier@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO supplier \\(id, name, address, phone, email, created_by, updated_by\\)").
			WithArgs(supplier.ID, supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", "admin").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(supplier)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("INSERT INTO supplier").
			WithArgs(supplier.ID, supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", "admin").
			WillReturnError(expectedErr)

		err := repo.Create(supplier)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestSupplierGet(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSupplierRepository(db)

	supplierID := "supp123"
	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow(supplierID, "Test Supplier", "123 Test St", "+1234567890", "supplier@example.com", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by FROM supplier WHERE id = \\$1").
			WithArgs(supplierID).
			WillReturnRows(rows)

		supplier, err := repo.Get(supplierID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if supplier == nil {
			t.Fatal("Expected supplier, got nil")
		}

		if supplier.ID != supplierID {
			t.Errorf("Expected ID %s, got %s", supplierID, supplier.ID)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		expectedErr := errors.New("no rows in result set")
		mock.ExpectQuery("SELECT id, name, address, phone, email, created_at, created_by, updated_at, updated_by FROM supplier WHERE id = \\$1").
			WithArgs(supplierID).
			WillReturnError(expectedErr)

		supplier, err := repo.Get(supplierID)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if supplier != nil && supplier.ID != "" {
			t.Errorf("Expected nil or empty supplier, got %v", supplier)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestSupplierUpdate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSupplierRepository(db)

	supplierID := "supp123"
	supplier := &model.Supplier{
		ID:      supplierID,
		Name:    "Updated Supplier",
		Address: "456 New St",
		Phone:   "+9876543210",
		Email:   "updated@example.com",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("UPDATE supplier SET name = \\$1, address = \\$2, phone = \\$3, email = \\$4, updated_by = \\$5, updated_at = now\\(\\) WHERE id = \\$6").
			WithArgs(supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", supplierID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(supplierID, supplier)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("UPDATE supplier SET").
			WithArgs(supplier.Name, supplier.Address, supplier.Phone, supplier.Email, "admin", supplierID).
			WillReturnError(expectedErr)

		err := repo.Update(supplierID, supplier)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestSupplierDelete(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSupplierRepository(db)

	supplierID := "supp123"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM supplier WHERE id = \\$1").
			WithArgs(supplierID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Delete(supplierID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("DELETE FROM supplier WHERE id = \\$1").
			WithArgs(supplierID).
			WillReturnError(expectedErr)

		err := repo.Delete(supplierID)
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestSupplierGetAll(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSupplierRepository(db)

	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success With Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow("supp123", "Supplier 1", "123 Test St", "+1234567890", "supp1@example.com", createdAt, "admin", updatedAt, "admin").
			AddRow("supp456", "Supplier 2", "456 Test St", "+0987654321", "supp2@example.com", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT \\* FROM supplier").
			WillReturnRows(rows)

		suppliers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(suppliers) != 2 {
			t.Errorf("Expected 2 suppliers, got %d", len(suppliers))
		}

		if suppliers[0].ID != "supp123" || suppliers[1].ID != "supp456" {
			t.Errorf("Unexpected supplier IDs: %v, %v", suppliers[0].ID, suppliers[1].ID)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Success With No Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address", "phone", "email", "created_at", "created_by", "updated_at", "updated_by"})

		mock.ExpectQuery("SELECT \\* FROM supplier").
			WillReturnRows(rows)

		suppliers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if suppliers == nil || len(suppliers) == 0 {
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
		mock.ExpectQuery("SELECT \\* FROM supplier").
			WillReturnError(expectedErr)

		suppliers, err := repo.GetAll()
		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if suppliers != nil {
			t.Errorf("Expected nil result, got %v", suppliers)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}
