package repository

import (
	"database/sql" // Import added
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GoodsChain/backend/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

// Helper to create a new mock DB and repository for tests
func newMockCarRepo(t *testing.T) (CarRepository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return NewCarRepository(sqlxDB), mock
}

// AnyTime argument for sqlmock
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}


func TestCarRepository_CreateCar(t *testing.T) {
	repo, mock := newMockCarRepo(t)
	carID := uuid.New().String()
	testCar := &model.Car{
		ID:         carID,
		Name:       "Test Car",
		SupplierID: uuid.New().String(),
		Price:      20000,
		CreatedBy:  "test_user",
		UpdatedBy:  "test_user",
	}

	query := regexp.QuoteMeta(`INSERT INTO car (id, name, supp_id, price, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)

	mock.ExpectExec(query).
		WithArgs(testCar.ID, testCar.Name, testCar.SupplierID, testCar.Price, AnyTime{}, testCar.CreatedBy, AnyTime{}, testCar.UpdatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateCar(testCar)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCarRepository_GetCarByID(t *testing.T) {
	repo, mock := newMockCarRepo(t)
	carID := uuid.New().String()
	expectedCar := &model.Car{
		ID:         carID,
		Name:       "Test Car",
		SupplierID: uuid.New().String(),
		Price:      20000,
		CreatedAt:  time.Now(),
		CreatedBy:  "test_user",
		UpdatedAt:  time.Now(),
		UpdatedBy:  "test_user",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "supp_id", "price", "created_at", "created_by", "updated_at", "updated_by"}).
		AddRow(expectedCar.ID, expectedCar.Name, expectedCar.SupplierID, expectedCar.Price, expectedCar.CreatedAt, expectedCar.CreatedBy, expectedCar.UpdatedAt, expectedCar.UpdatedBy)

	query := regexp.QuoteMeta(`SELECT id, name, supp_id, price, created_at, created_by, updated_at, updated_by FROM car WHERE id = $1`)
	mock.ExpectQuery(query).WithArgs(carID).WillReturnRows(rows)

	car, err := repo.GetCarByID(carID)
	assert.NoError(t, err)
	assert.NotNil(t, car)
	assert.Equal(t, expectedCar.ID, car.ID)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Test Not Found
	notFoundID := uuid.New().String()
	mock.ExpectQuery(query).WithArgs(notFoundID).WillReturnError(sql.ErrNoRows)
	car, err = repo.GetCarByID(notFoundID)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, car)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCarRepository_GetAllCars(t *testing.T) {
	repo, mock := newMockCarRepo(t)
	car1 := model.Car{ID: uuid.New().String(), Name: "Car 1", SupplierID: uuid.New().String(), Price: 100}
	car2 := model.Car{ID: uuid.New().String(), Name: "Car 2", SupplierID: uuid.New().String(), Price: 200}

	rows := sqlmock.NewRows([]string{"id", "name", "supp_id", "price", "created_at", "created_by", "updated_at", "updated_by"}).
		AddRow(car1.ID, car1.Name, car1.SupplierID, car1.Price, time.Now(), "user", time.Now(), "user").
		AddRow(car2.ID, car2.Name, car2.SupplierID, car2.Price, time.Now(), "user", time.Now(), "user")

	query := regexp.QuoteMeta(`SELECT id, name, supp_id, price, created_at, created_by, updated_at, updated_by FROM car ORDER BY created_at DESC`)
	mock.ExpectQuery(query).WillReturnRows(rows)

	cars, err := repo.GetAllCars()
	assert.NoError(t, err)
	assert.Len(t, cars, 2)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Test empty result
	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "supp_id", "price", "created_at", "created_by", "updated_at", "updated_by"}))
	cars, err = repo.GetAllCars()
	assert.NoError(t, err)
	assert.Len(t, cars, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCarRepository_UpdateCar(t *testing.T) {
	repo, mock := newMockCarRepo(t)
	carID := uuid.New().String()
	updatedCar := &model.Car{
		Name:       "Updated Test Car",
		SupplierID: uuid.New().String(),
		Price:      25000,
		UpdatedBy:  "updater_user",
	}

	query := regexp.QuoteMeta(`UPDATE car SET name = $1, supp_id = $2, price = $3, updated_at = $4, updated_by = $5 WHERE id = $6`)
	mock.ExpectExec(query).
		WithArgs(updatedCar.Name, updatedCar.SupplierID, updatedCar.Price, AnyTime{}, updatedCar.UpdatedBy, carID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 0 for lastInsertId, 1 for rowsAffected

	err := repo.UpdateCar(carID, updatedCar)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Test Not Found
	notFoundID := uuid.New().String()
	mock.ExpectExec(query).
		WithArgs(updatedCar.Name, updatedCar.SupplierID, updatedCar.Price, AnyTime{}, updatedCar.UpdatedBy, notFoundID).
		WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected
	err = repo.UpdateCar(notFoundID, updatedCar)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCarRepository_DeleteCar(t *testing.T) {
	repo, mock := newMockCarRepo(t)
	carID := uuid.New().String()

	query := regexp.QuoteMeta(`DELETE FROM car WHERE id = $1`)
	mock.ExpectExec(query).WithArgs(carID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteCar(carID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Test Not Found
	notFoundID := uuid.New().String()
	mock.ExpectExec(query).WithArgs(notFoundID).WillReturnResult(sqlmock.NewResult(0, 0))
	err = repo.DeleteCar(notFoundID)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.NoError(t, mock.ExpectationsWereMet())
}
