package service

import (
	"fmt"
	"inventory-api/internal/dto"
	"inventory-api/internal/model"
	"inventory-api/internal/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
	supplierRepo *repository.SupplierRepository
}

func NewProductService(productRepo *repository.ProductRepository, categoryRepo *repository.CategoryRepository, supplierRepo *repository.SupplierRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		categoryRepo: categoryRepo,
		supplierRepo: supplierRepo,
	}
}

func (s *ProductService) CreateProduct(req *dto.ProductRequest) error {
	_, err := s.productRepo.FindBySKU(req.SKU)
	if err == nil {
		return fmt.Errorf("Product with SKU %s already exists", req.SKU)
	}

	err = s.validateCategoryAndSupplier(req.CategoryID, req.SupplierID)
	if err != nil {
		return err
	}

	newProduct := &model.Product{
		SKU:         req.SKU,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		SupplierID:  req.SupplierID,
	}

	return s.productRepo.Create(newProduct)
}

func (s *ProductService) FindAllProducts() ([]model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *ProductService) FindProductByID(id uint) (*model.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *ProductService) UpdateProduct(id uint, req *dto.ProductRequest) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.validateCategoryAndSupplier(req.CategoryID, req.SupplierID)
	if err != nil {
		return err
	}

	if product.SKU != req.SKU {
		existingProduct, err := s.productRepo.FindBySKU(req.SKU)
		if err == nil && existingProduct.ID != id {
			return fmt.Errorf("Product with SKU %s already exists", req.SKU)
		}
	}

	product.SKU = req.SKU
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.SupplierID = req.SupplierID
	
	product.Category = model.Category{}
	product.Supplier = model.Supplier{}

	return s.productRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(id)
}

func (s *ProductService) validateCategoryAndSupplier(categoryID uint, supplierID uint) error {
	_, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return fmt.Errorf("Category with ID %d not found", categoryID)
	}

	_, err = s.supplierRepo.FindByID(supplierID)
	if err != nil {
		return fmt.Errorf("Supplier with ID %d not found", supplierID)
	}

	return nil
}