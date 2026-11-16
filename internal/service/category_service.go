package service

import (
	"inventory-api/internal/dto"
	"inventory-api/internal/model"
	"inventory-api/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(r *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: r}
}

func (s *CategoryService) CreateCategory(req dto.CategoryRequest) error {
	newCategory := &model.Category{
		Name: req.Name,
		Description: req.Description,
	}

	return s.categoryRepo.Create(newCategory)
}

func (s *CategoryService) FindAllCategories() ([]model.Category, error) {
	return s.categoryRepo.FindAll()
}

func (s *CategoryService) FindCategoryByID(id uint) (*model.Category, error) {
	return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) UpdateCategory(id uint, req *dto.CategoryRequest) error {
	updateCategory, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	updateCategory.Name = req.Name
	updateCategory.Description = req.Description

	return s.categoryRepo.Update(updateCategory)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.categoryRepo.Delete(id)
}