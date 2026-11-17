package repository

import (
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{db: database}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product

	err := r.db.Preload("Category").Preload("Supplier").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id uint) (*model.Product, error){
	var product model.Product

	err := r.db.Preload("Category").Preload("Supplier").First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindBySKU(sku string) (*model.Product, error){
	var product model.Product

	err := r.db.Preload("Category").Preload("Supplier").Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}