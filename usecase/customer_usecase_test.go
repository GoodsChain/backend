package usecase

import (
	"errors"
	"testing"

	"github.com/GoodsChain/backend/model"
	mock_repository "github.com/GoodsChain/backend/mock"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	usecase := NewCustomerUsecase(mockRepo)
	
	customer := &model.Customer{
		ID:   "1",
		Name: "Test Customer",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Create(customer).Return(nil)
		
		err := usecase.CreateCustomer(customer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().Create(customer).Return(expectedErr)
		
		err := usecase.CreateCustomer(customer)
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestGetCustomer(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	usecase := NewCustomerUsecase(mockRepo)
	
	customer := &model.Customer{
		ID:   "1",
		Name: "Test Customer",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Get("1").Return(customer, nil)
		
		result, err := usecase.GetCustomer("1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != customer {
			t.Errorf("Expected %v, got %v", customer, result)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("not found")
		mockRepo.EXPECT().Get("1").Return(nil, expectedErr)
		
		result, err := usecase.GetCustomer("1")
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}

func TestUpdateCustomer(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	usecase := NewCustomerUsecase(mockRepo)
	
	customer := &model.Customer{
		ID:   "1",
		Name: "Updated Customer",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Update("1", customer).Return(nil)
		
		err := usecase.UpdateCustomer("1", customer)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		mockRepo.EXPECT().Update("1", customer).Return(expectedErr)
		
		err := usecase.UpdateCustomer("1", customer)
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestDeleteCustomer(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	usecase := NewCustomerUsecase(mockRepo)
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Delete("1").Return(nil)
		
		err := usecase.DeleteCustomer("1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("delete error")
		mockRepo.EXPECT().Delete("1").Return(expectedErr)
		
		err := usecase.DeleteCustomer("1")
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestGetAllCustomers(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	usecase := NewCustomerUsecase(mockRepo)
	
	customers := []*model.Customer{
		{ID: "1", Name: "Customer 1"},
		{ID: "2", Name: "Customer 2"},
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetAll().Return(customers, nil)
		
		result, err := usecase.GetAllCustomers()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != len(customers) {
			t.Errorf("Expected %d customers, got %d", len(customers), len(result))
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().GetAll().Return(nil, expectedErr)
		
		result, err := usecase.GetAllCustomers()
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}
