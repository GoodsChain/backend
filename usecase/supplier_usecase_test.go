package usecase

import (
	"errors"
	"testing"

	"github.com/GoodsChain/backend/model"
	mock_repository "github.com/GoodsChain/backend/mock"
	"go.uber.org/mock/gomock"
)

func TestCreateSupplier(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockSupplierRepository(ctrl)
	usecase := NewSupplierUsecase(mockRepo)
	
	supplier := &model.Supplier{
		ID:   "1",
		Name: "Test Supplier",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Create(supplier).Return(nil)
		
		err := usecase.CreateSupplier(supplier)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().Create(supplier).Return(expectedErr)
		
		err := usecase.CreateSupplier(supplier)
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestGetSupplier(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockSupplierRepository(ctrl)
	usecase := NewSupplierUsecase(mockRepo)
	
	supplier := &model.Supplier{
		ID:   "1",
		Name: "Test Supplier",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Get("1").Return(supplier, nil)
		
		result, err := usecase.GetSupplier("1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != supplier {
			t.Errorf("Expected %v, got %v", supplier, result)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("not found")
		mockRepo.EXPECT().Get("1").Return(nil, expectedErr)
		
		result, err := usecase.GetSupplier("1")
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}

func TestUpdateSupplier(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockSupplierRepository(ctrl)
	usecase := NewSupplierUsecase(mockRepo)
	
	supplier := &model.Supplier{
		ID:   "1",
		Name: "Updated Supplier",
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Update("1", supplier).Return(nil)
		
		err := usecase.UpdateSupplier("1", supplier)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		mockRepo.EXPECT().Update("1", supplier).Return(expectedErr)
		
		err := usecase.UpdateSupplier("1", supplier)
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestDeleteSupplier(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockSupplierRepository(ctrl)
	usecase := NewSupplierUsecase(mockRepo)
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().Delete("1").Return(nil)
		
		err := usecase.DeleteSupplier("1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("delete error")
		mockRepo.EXPECT().Delete("1").Return(expectedErr)
		
		err := usecase.DeleteSupplier("1")
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
	})
}

func TestGetAllSuppliers(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mockRepo := mock_repository.NewMockSupplierRepository(ctrl)
	usecase := NewSupplierUsecase(mockRepo)
	
	suppliers := []*model.Supplier{
		{ID: "1", Name: "Supplier 1"},
		{ID: "2", Name: "Supplier 2"},
	}
	
	// Test cases
	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetAll().Return(suppliers, nil)
		
		result, err := usecase.GetAllSuppliers()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != len(suppliers) {
			t.Errorf("Expected %d suppliers, got %d", len(suppliers), len(result))
		}
	})
	
	t.Run("Repository Error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.EXPECT().GetAll().Return(nil, expectedErr)
		
		result, err := usecase.GetAllSuppliers()
		if err != expectedErr {
			t.Errorf("Expected %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}
