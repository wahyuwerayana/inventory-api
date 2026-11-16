package service

import (
	"inventory-api/internal/dto"
	"inventory-api/internal/model"
	"inventory-api/internal/repository"
)

type SupplierService struct {
	supplierRepo *repository.SupplierRepository
}

func NewSupplierService(r *repository.SupplierRepository) *SupplierService {
	return &SupplierService{supplierRepo: r}
}

func (s *SupplierService) CreateSupplier(req *dto.SupplierRequest) error {
	newSupplier := &model.Supplier{
		Name:  req.Name,
		Contact: req.Contact,
		Phone: req.Phone,
		Address: req.Address,
	}

	return s.supplierRepo.Create(newSupplier)
}

func (s *SupplierService) FindAllSuppliers() ([]model.Supplier, error) {
	return s.supplierRepo.FindAll()
}

func (s *SupplierService) FindSupplierByID(id uint) (*model.Supplier, error) {
	return s.supplierRepo.FindByID(id)
}

func (s *SupplierService) UpdateSupplier(id uint, req *dto.SupplierRequest) error {
	supplier, err := s.FindSupplierByID(id)
	if err != nil {
		return err
	}

	supplier.Name = req.Name
	supplier.Contact = req.Contact
	supplier.Phone = req.Phone
	supplier.Address = req.Address

	return s.supplierRepo.Update(supplier)
}

func (s *SupplierService) DeleteSupplier(id uint) error {
	_, err := s.FindSupplierByID(id)
	if err != nil {
		return err
	}

	return s.supplierRepo.Delete(id)
}