package usecase

import (
	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/repository"
	"github.com/google/uuid"
	// "time" // Not strictly needed here if repository handles timestamps
)

// CarUsecase defines the interface for car business logic
type CarUsecase interface {
	CreateCar(car *model.Car) error
	GetCar(id string) (*model.Car, error)
	GetAllCars() ([]model.Car, error)
	UpdateCar(id string, car *model.Car) error
	DeleteCar(id string) error
}

type carUsecase struct {
	carRepo repository.CarRepository
}

// NewCarUsecase creates a new instance of CarUsecase
func NewCarUsecase(carRepo repository.CarRepository) CarUsecase {
	return &carUsecase{carRepo: carRepo}
}

// CreateCar handles the business logic for creating a new car
func (uc *carUsecase) CreateCar(car *model.Car) error {
	if car.ID == "" {
		car.ID = uuid.New().String()
	}
	// Assuming CreatedBy/UpdatedBy are set by a higher layer (e.g., handler from auth context)
	// or a default system user if not provided. For now, let's assume they might be pre-filled
	// or should be explicitly passed to the usecase.
	// car.CreatedAt = time.Now() // Repository handles this
	// car.UpdatedAt = time.Now() // Repository handles this

	// Example: Set default user if not provided, this logic can be more sophisticated
	if car.CreatedBy == "" {
		car.CreatedBy = "system" // Or get from context
	}
	if car.UpdatedBy == "" {
		car.UpdatedBy = "system" // Or get from context
	}

	return uc.carRepo.CreateCar(car)
}

// GetCar retrieves a car by its ID
func (uc *carUsecase) GetCar(id string) (*model.Car, error) {
	return uc.carRepo.GetCarByID(id)
}

// GetAllCars retrieves all cars
func (uc *carUsecase) GetAllCars() ([]model.Car, error) {
	return uc.carRepo.GetAllCars()
}

// UpdateCar handles the business logic for updating an existing car
func (uc *carUsecase) UpdateCar(id string, car *model.Car) error {
	// Ensure UpdatedBy is set
	if car.UpdatedBy == "" {
		car.UpdatedBy = "system" // Or get from context
	}
	// car.UpdatedAt = time.Now() // Repository handles this

	// Optional: Could fetch existing car to ensure it exists before update,
	// or to merge fields if partial updates are allowed.
	// For now, repository's UpdateCar handles not-found error.
	return uc.carRepo.UpdateCar(id, car)
}

// DeleteCar handles the business logic for deleting a car
func (uc *carUsecase) DeleteCar(id string) error {
	return uc.carRepo.DeleteCar(id)
}
