package repository

import (
	"fmt"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type StockMovementRepository struct {
	db *gorm.DB
}

func NewStockMovementRepository(database *gorm.DB) *StockMovementRepository {
	return &StockMovementRepository{db: database}
}

func (r *StockMovementRepository) Create(stockMovement *model.StockMovement) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var product model.Product
		
		if err := tx.First(&product, stockMovement.ProductID).Error; err != nil {
			return err
		}

		switch stockMovement.Type {
			case "IN":
				product.Stock += stockMovement.Quantity
			case "OUT":
				if product.Stock < stockMovement.Quantity {
					return fmt.Errorf("Insufficient Stock")
				}
				product.Stock -= stockMovement.Quantity
		}

		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		return tx.Create(stockMovement).Error
	})
}

func (r *StockMovementRepository) FindAll() ([]model.StockMovement, error) {
	var stockMovements []model.StockMovement

	err := r.db.Preload("User").
			Preload("Product.Category").
			Preload("Product.Supplier").
			Find(&stockMovements).Error
	if err != nil {
		return nil, err
	}

	return stockMovements, nil
}

func (r *StockMovementRepository) FindByID(id uint) (*model.StockMovement, error) {
	var stockMovement model.StockMovement
	
	err := r.db.Preload("User").
		Preload("Product.Category").
		Preload("Product.Supplier").
		First(&stockMovement, id).Error
	if err != nil {
		return nil, err
	}

	return &stockMovement, nil
}