package repository

import (
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: database}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	
	err := r.db.Find(&categories).Error

	if(err != nil){
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error){
	var category model.Category

	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}