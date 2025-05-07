package usecase

import (
	"errors"
	"testing"

	"github.com/GoodsChain/backend/model"
	mock_repository "github.com/GoodsChain/backend/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomerCar(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerCar := &model.CustomerCar{
		ID:         "",
		CarID:      "car123",
		CustomerID: "cust123",
	}
	
	t.Run("Success With ID Generation", func(t *testing.T) {
		mockRepo.EXPECT().
			Create(gomock.Any()).
			DoAndReturn(func(cc *model.CustomerCar) error {
				assert.NotEmpty(t, cc.ID)
				assert.Equal(t, "car123", cc.CarID)
				assert.Equal(t, "cust123", cc.CustomerID)
				assert.Equal(t, "system", cc.CreatedBy)
				assert.Equal(t, "system", cc.UpdatedBy)
				return nil
			})
		
		err := usecase.CreateCustomerCar(customerCar)
		assert.NoError(t, err)
	})
	
	t.Run("Success With Provided ID", func(t *testing.T) {
		customerCarWithID := &model.CustomerCar{
			ID:         "cc123",
			CarID:      "car123",
			CustomerID: "cust123",
		}
		
		mockRepo.EXPECT().
			Create(gomock.Any()).
			DoAndReturn(func(cc *model.CustomerCar) error {
				assert.Equal(t, "cc123", cc.ID)
				return nil
			})
		
		err := usecase.CreateCustomerCar(customerCarWithID)
		assert.NoError(t, err)
	})
	
	t.Run("Success With Provided CreatedBy", func(t *testing.T) {
		customerCarWithCreator := &model.CustomerCar{
			ID:         "cc123",
			CarID:      "car123",
			CustomerID: "cust123",
			CreatedBy:  "admin",
		}
		
		mockRepo.EXPECT().
			Create(gomock.Any()).
			DoAndReturn(func(cc *model.CustomerCar) error {
				assert.Equal(t, "admin", cc.CreatedBy)
				assert.Equal(t, "admin", cc.UpdatedBy)
				return nil
			})
		
		err := usecase.CreateCustomerCar(customerCarWithCreator)
		assert.NoError(t, err)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().Create(gomock.Any()).Return(expectedErr)
		
		err := usecase.CreateCustomerCar(customerCar)
		assert.Equal(t, expectedErr, err)
	})
}

func TestGetCustomerCar(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerCar := &model.CustomerCar{
		ID:         "cc123",
		CarID:      "car123",
		CustomerID: "cust123",
	}
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetByID("cc123").Return(customerCar, nil)
		
		result, err := usecase.GetCustomerCar("cc123")
		assert.NoError(t, err)
		assert.Equal(t, customerCar, result)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("not found")
		mockRepo.EXPECT().GetByID("cc123").Return(nil, expectedErr)
		
		result, err := usecase.GetCustomerCar("cc123")
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
	})
}

func TestGetAllCustomerCars(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: "car123", CustomerID: "cust123"},
		{ID: "cc456", CarID: "car456", CustomerID: "cust456"},
	}
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetAll().Return(customerCars, nil)
		
		result, err := usecase.GetAllCustomerCars()
		assert.NoError(t, err)
		assert.Equal(t, customerCars, result)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().GetAll().Return(nil, expectedErr)
		
		result, err := usecase.GetAllCustomerCars()
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
	})
}

func TestGetCustomerCarsByCustomerID(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerID := "cust123"
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: "car123", CustomerID: customerID},
		{ID: "cc456", CarID: "car456", CustomerID: customerID},
	}
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetByCustomerID(customerID).Return(customerCars, nil)
		
		result, err := usecase.GetCustomerCarsByCustomerID(customerID)
		assert.NoError(t, err)
		assert.Equal(t, customerCars, result)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().GetByCustomerID(customerID).Return(nil, expectedErr)
		
		result, err := usecase.GetCustomerCarsByCustomerID(customerID)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
	})
}

func TestGetCustomerCarsByCarID(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	carID := "car123"
	customerCars := []*model.CustomerCar{
		{ID: "cc123", CarID: carID, CustomerID: "cust123"},
		{ID: "cc456", CarID: carID, CustomerID: "cust456"},
	}
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetByCarID(carID).Return(customerCars, nil)
		
		result, err := usecase.GetCustomerCarsByCarID(carID)
		assert.NoError(t, err)
		assert.Equal(t, customerCars, result)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().GetByCarID(carID).Return(nil, expectedErr)
		
		result, err := usecase.GetCustomerCarsByCarID(carID)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
	})
}

func TestUpdateCustomerCar(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerCarID := "cc123"
	customerCar := &model.CustomerCar{
		ID:         customerCarID,
		CarID:      "car456",
		CustomerID: "cust456",
	}
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().
			Update(customerCarID, gomock.Any()).
			DoAndReturn(func(id string, cc *model.CustomerCar) error {
				assert.Equal(t, "system", cc.UpdatedBy)
				return nil
			})
		
		err := usecase.UpdateCustomerCar(customerCarID, customerCar)
		assert.NoError(t, err)
	})
	
	t.Run("Success with Provided UpdatedBy", func(t *testing.T) {
		customerCarWithUpdater := &model.CustomerCar{
			ID:         customerCarID,
			CarID:      "car456",
			CustomerID: "cust456",
			UpdatedBy:  "admin",
		}
		
		mockRepo.EXPECT().
			Update(customerCarID, gomock.Any()).
			DoAndReturn(func(id string, cc *model.CustomerCar) error {
				assert.Equal(t, "admin", cc.UpdatedBy)
				return nil
			})
		
		err := usecase.UpdateCustomerCar(customerCarID, customerCarWithUpdater)
		assert.NoError(t, err)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		mockRepo.EXPECT().Update(customerCarID, gomock.Any()).Return(expectedErr)
		
		err := usecase.UpdateCustomerCar(customerCarID, customerCar)
		assert.Equal(t, expectedErr, err)
	})
}

func TestDeleteCustomerCar(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerCarRepository(ctrl)
	usecase := NewCustomerCarUsecase(mockRepo)
	
	customerCarID := "cc123"
	
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Delete(customerCarID).Return(nil)
		
		err := usecase.DeleteCustomerCar(customerCarID)
		assert.NoError(t, err)
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("delete error")
		mockRepo.EXPECT().Delete(customerCarID).Return(expectedErr)
		
		err := usecase.DeleteCustomerCar(customerCarID)
		assert.Equal(t, expectedErr, err)
	})
}
