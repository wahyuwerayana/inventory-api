package repository

import (
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(database *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db: database}
}

func (r *SupplierRepository) Create(supplier *model.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *SupplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	
	err := r.db.Find(&suppliers).Error
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r *SupplierRepository) FindByID(id uint) (*model.Supplier, error){
	var supplier model.Supplier

	err := r.db.First(&supplier, id).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *SupplierRepository) Update(supplier *model.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *SupplierRepository) Delete(id uint) error {
	return r.db.Delete(&model.Supplier{}, id).Error
}