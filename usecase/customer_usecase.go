package usecase

import (
	"github.com/GoodsChain/backend/repository"
	"github.com/GoodsChain/backend/model"
)

type CustomerUsecase interface {
	CreateCustomer(customer *model.Customer) error
	GetCustomer(id string) (*model.Customer, error)
	UpdateCustomer(id string, customer *model.Customer) error
	DeleteCustomer(id string) error
	GetAllCustomers() ([]*model.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(customerRepo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (u *customerUsecase) CreateCustomer(customer *model.Customer) error {
	return u.customerRepo.Create(customer)
}

func (u *customerUsecase) GetCustomer(id string) (*model.Customer, error) {
	return u.customerRepo.Get(id)
}

func (u *customerUsecase) UpdateCustomer(id string, customer *model.Customer) error {
	return u.customerRepo.Update(id, customer)
}

func (u *customerUsecase) DeleteCustomer(id string) error {
	return u.customerRepo.Delete(id)
}

func (u *customerUsecase) GetAllCustomers() ([]*model.Customer, error) {
	return u.customerRepo.GetAll()
}
