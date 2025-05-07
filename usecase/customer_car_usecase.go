package usecase

import (
	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/repository"
	"github.com/google/uuid"
)

// CustomerCarUsecase defines the interface for customer car business logic
type CustomerCarUsecase interface {
	CreateCustomerCar(customerCar *model.CustomerCar) error
	GetCustomerCar(id string) (*model.CustomerCar, error)
	GetAllCustomerCars() ([]*model.CustomerCar, error)
	GetCustomerCarsByCustomerID(customerID string) ([]*model.CustomerCar, error)
	GetCustomerCarsByCarID(carID string) ([]*model.CustomerCar, error)
	UpdateCustomerCar(id string, customerCar *model.CustomerCar) error
	DeleteCustomerCar(id string) error
}

type customerCarUsecase struct {
	customerCarRepo repository.CustomerCarRepository
}

// NewCustomerCarUsecase creates a new instance of CustomerCarUsecase
func NewCustomerCarUsecase(customerCarRepo repository.CustomerCarRepository) CustomerCarUsecase {
	return &customerCarUsecase{
		customerCarRepo: customerCarRepo,
	}
}

// CreateCustomerCar handles the business logic for creating a new customer car relationship
func (u *customerCarUsecase) CreateCustomerCar(customerCar *model.CustomerCar) error {
	// Generate UUID if not provided
	if customerCar.ID == "" {
		customerCar.ID = uuid.New().String()
	}

	// Set default values for created_by and updated_by if not provided
	if customerCar.CreatedBy == "" {
		customerCar.CreatedBy = "system"
	}
	if customerCar.UpdatedBy == "" {
		customerCar.UpdatedBy = customerCar.CreatedBy
	}

	return u.customerCarRepo.Create(customerCar)
}

// GetCustomerCar retrieves a customer car relationship by ID
func (u *customerCarUsecase) GetCustomerCar(id string) (*model.CustomerCar, error) {
	return u.customerCarRepo.GetByID(id)
}

// GetAllCustomerCars retrieves all customer car relationships
func (u *customerCarUsecase) GetAllCustomerCars() ([]*model.CustomerCar, error) {
	return u.customerCarRepo.GetAll()
}

// GetCustomerCarsByCustomerID retrieves all car relationships for a specific customer
func (u *customerCarUsecase) GetCustomerCarsByCustomerID(customerID string) ([]*model.CustomerCar, error) {
	return u.customerCarRepo.GetByCustomerID(customerID)
}

// GetCustomerCarsByCarID retrieves all customer relationships for a specific car
func (u *customerCarUsecase) GetCustomerCarsByCarID(carID string) ([]*model.CustomerCar, error) {
	return u.customerCarRepo.GetByCarID(carID)
}

// UpdateCustomerCar updates an existing customer car relationship
func (u *customerCarUsecase) UpdateCustomerCar(id string, customerCar *model.CustomerCar) error {
	// Set default value for updated_by if not provided
	if customerCar.UpdatedBy == "" {
		customerCar.UpdatedBy = "system"
	}

	return u.customerCarRepo.Update(id, customerCar)
}

// DeleteCustomerCar removes a customer car relationship by ID
func (u *customerCarUsecase) DeleteCustomerCar(id string) error {
	return u.customerCarRepo.Delete(id)
}
