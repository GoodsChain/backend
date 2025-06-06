package usecase

import (
	"errors"
	"testing"
	// "time" // Not needed for these tests as repo mock handles time

	"github.com/GoodsChain/backend/mock" // Assuming mock package is at this path
	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/repository" // For repository.ErrNotFound
	"go.uber.org/mock/gomock"                 // Corrected import path
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCarUsecase_CreateCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarRepo := mock.NewMockCarRepository(ctrl)
	uc := NewCarUsecase(mockCarRepo)

	car := &model.Car{Name: "Test Car", SupplierID: "supp1", Price: 10000}
	expectedCar := *car
	// ID will be generated by usecase if empty
	// CreatedBy/UpdatedBy will be defaulted by usecase if empty

	// Test case 1: Successful creation
	mockCarRepo.EXPECT().CreateCar(gomock.Any()).DoAndReturn(
		func(c *model.Car) error {
			assert.NotEmpty(t, c.ID)
			assert.Equal(t, expectedCar.Name, c.Name)
			assert.Equal(t, "system", c.CreatedBy) // Default value
			assert.Equal(t, "system", c.UpdatedBy) // Default value
			return nil
		}).Times(1)

	err := uc.CreateCar(car)
	assert.NoError(t, err)

	// Test case 2: Repository returns an error
	repoErr := errors.New("repository error")
	mockCarRepo.EXPECT().CreateCar(gomock.Any()).Return(repoErr).Times(1)

	carWithID := &model.Car{ID: uuid.New().String(), Name: "Test Car 2", CreatedBy: "user1", UpdatedBy: "user1"}
	err = uc.CreateCar(carWithID)
	assert.EqualError(t, err, "repository error")
}

func TestCarUsecase_GetCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarRepo := mock.NewMockCarRepository(ctrl)
	uc := NewCarUsecase(mockCarRepo)

	carID := uuid.New().String()
	expectedCar := &model.Car{ID: carID, Name: "Found Car"}

	// Test case 1: Successful retrieval
	mockCarRepo.EXPECT().GetCarByID(carID).Return(expectedCar, nil).Times(1)
	retrievedCar, err := uc.GetCar(carID)
	assert.NoError(t, err)
	assert.Equal(t, expectedCar, retrievedCar)

	// Test case 2: Car not found
	notFoundID := uuid.New().String()
	mockCarRepo.EXPECT().GetCarByID(notFoundID).Return(nil, repository.ErrNotFound).Times(1)
	retrievedCar, err = uc.GetCar(notFoundID)
	assert.ErrorIs(t, err, repository.ErrNotFound)
	assert.Nil(t, retrievedCar)

	// Test case 3: Other repository error
	errorID := uuid.New().String()
	repoErr := errors.New("some db error")
	mockCarRepo.EXPECT().GetCarByID(errorID).Return(nil, repoErr).Times(1)
	retrievedCar, err = uc.GetCar(errorID)
	assert.EqualError(t, err, "some db error")
	assert.Nil(t, retrievedCar)
}

func TestCarUsecase_GetAllCars(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarRepo := mock.NewMockCarRepository(ctrl)
	uc := NewCarUsecase(mockCarRepo)

	expectedCars := []model.Car{
		{ID: uuid.New().String(), Name: "Car 1"},
		{ID: uuid.New().String(), Name: "Car 2"},
	}

	// Test case 1: Successful retrieval
	mockCarRepo.EXPECT().GetAllCars().Return(expectedCars, nil).Times(1)
	cars, err := uc.GetAllCars()
	assert.NoError(t, err)
	assert.Equal(t, expectedCars, cars)

	// Test case 2: Empty list
	mockCarRepo.EXPECT().GetAllCars().Return([]model.Car{}, nil).Times(1)
	cars, err = uc.GetAllCars()
	assert.NoError(t, err)
	assert.Empty(t, cars)

	// Test case 3: Repository error
	repoErr := errors.New("db query failed")
	mockCarRepo.EXPECT().GetAllCars().Return(nil, repoErr).Times(1)
	cars, err = uc.GetAllCars()
	assert.EqualError(t, err, "db query failed")
	assert.Nil(t, cars)
}

func TestCarUsecase_UpdateCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarRepo := mock.NewMockCarRepository(ctrl)
	uc := NewCarUsecase(mockCarRepo)

	carID := uuid.New().String()
	carToUpdate := &model.Car{Name: "Updated Car Name"}

	// Test case 1: Successful update
	mockCarRepo.EXPECT().UpdateCar(carID, gomock.Any()).DoAndReturn(
		func(id string, c *model.Car) error {
			assert.Equal(t, carID, id)
			assert.Equal(t, carToUpdate.Name, c.Name)
			assert.Equal(t, "system", c.UpdatedBy) // Default value
			return nil
		}).Times(1)
	err := uc.UpdateCar(carID, carToUpdate)
	assert.NoError(t, err)

	// Test case 2: Car not found by repository
	notFoundID := uuid.New().String()
	mockCarRepo.EXPECT().UpdateCar(notFoundID, gomock.Any()).Return(repository.ErrNotFound).Times(1)
	err = uc.UpdateCar(notFoundID, carToUpdate)
	assert.ErrorIs(t, err, repository.ErrNotFound)

	// Test case 3: Other repository error
	errorID := uuid.New().String()
	repoErr := errors.New("update failed")
	carWithUser := &model.Car{Name: "Updated Car Name", UpdatedBy: "user1"}
	mockCarRepo.EXPECT().UpdateCar(errorID, carWithUser).Return(repoErr).Times(1)
	err = uc.UpdateCar(errorID, carWithUser)
	assert.EqualError(t, err, "update failed")
}

func TestCarUsecase_DeleteCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarRepo := mock.NewMockCarRepository(ctrl)
	uc := NewCarUsecase(mockCarRepo)

	carID := uuid.New().String()

	// Test case 1: Successful deletion
	mockCarRepo.EXPECT().DeleteCar(carID).Return(nil).Times(1)
	err := uc.DeleteCar(carID)
	assert.NoError(t, err)

	// Test case 2: Car not found by repository
	notFoundID := uuid.New().String()
	mockCarRepo.EXPECT().DeleteCar(notFoundID).Return(repository.ErrNotFound).Times(1)
	err = uc.DeleteCar(notFoundID)
	assert.ErrorIs(t, err, repository.ErrNotFound)

	// Test case 3: Other repository error
	errorID := uuid.New().String()
	repoErr := errors.New("delete failed")
	mockCarRepo.EXPECT().DeleteCar(errorID).Return(repoErr).Times(1)
	err = uc.DeleteCar(errorID)
	assert.EqualError(t, err, "delete failed")
}
