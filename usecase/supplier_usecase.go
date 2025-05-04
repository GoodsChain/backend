package usecase

import (
	"github.com/GoodsChain/backend/repository"
	"github.com/GoodsChain/backend/model"
)

type SupplierUsecase interface {
	CreateSupplier(supplier *model.Supplier) error
	GetSupplier(id string) (*model.Supplier, error)
	UpdateSupplier(id string, supplier *model.Supplier) error
	DeleteSupplier(id string) error
	GetAllSuppliers() ([]*model.Supplier, error)
}

type supplierUsecase struct {
	supplierRepo repository.SupplierRepository
}

func NewSupplierUsecase(supplierRepo repository.SupplierRepository) SupplierUsecase {
	return &supplierUsecase{
		supplierRepo: supplierRepo,
	}
}

func (u *supplierUsecase) CreateSupplier(supplier *model.Supplier) error {
	return u.supplierRepo.Create(supplier)
}

func (u *supplierUsecase) GetSupplier(id string) (*model.Supplier, error) {
	return u.supplierRepo.Get(id)
}

func (u *supplierUsecase) UpdateSupplier(id string, supplier *model.Supplier) error {
	return u.supplierRepo.Update(id, supplier)
}

func (u *supplierUsecase) DeleteSupplier(id string) error {
	return u.supplierRepo.Delete(id)
}

func (u *supplierUsecase) GetAllSuppliers() ([]*model.Supplier, error) {
	return u.supplierRepo.GetAll()
}
