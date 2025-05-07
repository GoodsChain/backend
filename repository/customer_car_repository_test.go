package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GoodsChain/backend/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomerCarRepository(t *testing.T) {
	db, _ := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	if repo == nil {
		t.Fatal("Expected repository to be created, got nil")
	}
}

func TestCustomerCarCreate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	customerCar := &model.CustomerCar{
		ID:         "cc123",
		CarID:      "car123",
		CustomerID: "cust123",
		CreatedBy:  "admin",
		UpdatedBy:  "admin",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO customer_car \\(id, car_id, cust_id, created_at, created_by, updated_at, updated_by\\)").
			WithArgs(customerCar.ID, customerCar.CarID, customerCar.CustomerID,
				sqlmock.AnyArg(), customerCar.CreatedBy, sqlmock.AnyArg(), customerCar.UpdatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(customerCar)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("INSERT INTO customer_car").
			WithArgs(customerCar.ID, customerCar.CarID, customerCar.CustomerID,
				sqlmock.AnyArg(), customerCar.CreatedBy, sqlmock.AnyArg(), customerCar.UpdatedBy).
			WillReturnError(expectedErr)

		err := repo.Create(customerCar)
		assert.Equal(t, expectedErr, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarGetByID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	customerCarID := "cc123"
	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow(customerCarID, "car123", "cust123", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnRows(rows)

		customerCar, err := repo.GetByID(customerCarID)
		assert.NoError(t, err)
		assert.NotNil(t, customerCar)
		assert.Equal(t, customerCarID, customerCar.ID)
		assert.Equal(t, "car123", customerCar.CarID)
		assert.Equal(t, "cust123", customerCar.CustomerID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnError(sql.ErrNoRows)

		customerCar, err := repo.GetByID(customerCarID)
		assert.Equal(t, ErrNotFound, err)
		assert.Nil(t, customerCar)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnError(expectedErr)

		customerCar, err := repo.GetByID(customerCarID)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customerCar)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarGetAll(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success With Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow("cc123", "car123", "cust123", createdAt, "admin", updatedAt, "admin").
			AddRow("cc456", "car456", "cust456", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car ORDER BY created_at DESC").
			WillReturnRows(rows)

		customerCars, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, customerCars, 2)
		assert.Equal(t, "cc123", customerCars[0].ID)
		assert.Equal(t, "cc456", customerCars[1].ID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Success With No Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"})

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car ORDER BY created_at DESC").
			WillReturnRows(rows)

		customerCars, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Empty(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car ORDER BY created_at DESC").
			WillReturnError(expectedErr)

		customerCars, err := repo.GetAll()
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarGetByCustomerID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	customerID := "cust123"
	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success With Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow("cc123", "car123", customerID, createdAt, "admin", updatedAt, "admin").
			AddRow("cc456", "car456", customerID, createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE cust_id = \\$1 ORDER BY created_at DESC").
			WithArgs(customerID).
			WillReturnRows(rows)

		customerCars, err := repo.GetByCustomerID(customerID)
		assert.NoError(t, err)
		assert.Len(t, customerCars, 2)
		assert.Equal(t, customerID, customerCars[0].CustomerID)
		assert.Equal(t, customerID, customerCars[1].CustomerID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Success With No Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"})

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE cust_id = \\$1 ORDER BY created_at DESC").
			WithArgs(customerID).
			WillReturnRows(rows)

		customerCars, err := repo.GetByCustomerID(customerID)
		assert.NoError(t, err)
		assert.Empty(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE cust_id = \\$1 ORDER BY created_at DESC").
			WithArgs(customerID).
			WillReturnError(expectedErr)

		customerCars, err := repo.GetByCustomerID(customerID)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarGetByCarID(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	carID := "car123"
	createdAt := time.Now()
	updatedAt := time.Now()

	t.Run("Success With Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"}).
			AddRow("cc123", carID, "cust123", createdAt, "admin", updatedAt, "admin").
			AddRow("cc456", carID, "cust456", createdAt, "admin", updatedAt, "admin")

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE car_id = \\$1 ORDER BY created_at DESC").
			WithArgs(carID).
			WillReturnRows(rows)

		customerCars, err := repo.GetByCarID(carID)
		assert.NoError(t, err)
		assert.Len(t, customerCars, 2)
		assert.Equal(t, carID, customerCars[0].CarID)
		assert.Equal(t, carID, customerCars[1].CarID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Success With No Results", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "car_id", "cust_id", "created_at", "created_by", "updated_at", "updated_by"})

		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE car_id = \\$1 ORDER BY created_at DESC").
			WithArgs(carID).
			WillReturnRows(rows)

		customerCars, err := repo.GetByCarID(carID)
		assert.NoError(t, err)
		assert.Empty(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectQuery("SELECT id, car_id, cust_id, created_at, created_by, updated_at, updated_by FROM customer_car WHERE car_id = \\$1 ORDER BY created_at DESC").
			WithArgs(carID).
			WillReturnError(expectedErr)

		customerCars, err := repo.GetByCarID(carID)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customerCars)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarUpdate(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	customerCarID := "cc123"
	customerCar := &model.CustomerCar{
		ID:         customerCarID,
		CarID:      "car456",
		CustomerID: "cust456",
		UpdatedBy:  "admin",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("UPDATE customer_car SET car_id = \\$1, cust_id = \\$2, updated_at = \\$3, updated_by = \\$4 WHERE id = \\$5").
			WithArgs(customerCar.CarID, customerCar.CustomerID, sqlmock.AnyArg(), customerCar.UpdatedBy, customerCarID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(customerCarID, customerCar)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectExec("UPDATE customer_car SET car_id = \\$1, cust_id = \\$2, updated_at = \\$3, updated_by = \\$4 WHERE id = \\$5").
			WithArgs(customerCar.CarID, customerCar.CustomerID, sqlmock.AnyArg(), customerCar.UpdatedBy, customerCarID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Update(customerCarID, customerCar)
		assert.Equal(t, ErrNotFound, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("UPDATE customer_car SET car_id = \\$1, cust_id = \\$2, updated_at = \\$3, updated_by = \\$4 WHERE id = \\$5").
			WithArgs(customerCar.CarID, customerCar.CustomerID, sqlmock.AnyArg(), customerCar.UpdatedBy, customerCarID).
			WillReturnError(expectedErr)

		err := repo.Update(customerCarID, customerCar)
		assert.Equal(t, expectedErr, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Result Error", func(t *testing.T) {
		expectedErr := errors.New("result error")
		mock.ExpectExec("UPDATE customer_car SET car_id = \\$1, cust_id = \\$2, updated_at = \\$3, updated_by = \\$4 WHERE id = \\$5").
			WithArgs(customerCar.CarID, customerCar.CustomerID, sqlmock.AnyArg(), customerCar.UpdatedBy, customerCarID).
			WillReturnResult(sqlmock.NewErrorResult(expectedErr))

		err := repo.Update(customerCarID, customerCar)
		assert.Equal(t, expectedErr, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}

func TestCustomerCarDelete(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewCustomerCarRepository(db)

	customerCarID := "cc123"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Delete(customerCarID)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete(customerCarID)
		assert.Equal(t, ErrNotFound, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mock.ExpectExec("DELETE FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnError(expectedErr)

		err := repo.Delete(customerCarID)
		assert.Equal(t, expectedErr, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})

	t.Run("Result Error", func(t *testing.T) {
		expectedErr := errors.New("result error")
		mock.ExpectExec("DELETE FROM customer_car WHERE id = \\$1").
			WithArgs(customerCarID).
			WillReturnResult(sqlmock.NewErrorResult(expectedErr))

		err := repo.Delete(customerCarID)
		assert.Equal(t, expectedErr, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}
