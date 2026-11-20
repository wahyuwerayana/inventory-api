package service

import (
	"inventory-api/internal/dto"
	"inventory-api/internal/model"
	"inventory-api/internal/repository"
)

type StockMovementService struct {
	stockMovementRepo *repository.StockMovementRepository
	productRepo       *repository.ProductRepository
}

func NewStockMovementService(stockMovementRepo *repository.StockMovementRepository, productRepo *repository.ProductRepository) *StockMovementService {
	return &StockMovementService{
		stockMovementRepo: stockMovementRepo,
		productRepo:       productRepo,
	}
}

func (s *StockMovementService) CreateStockMovement(userID uint, req *dto.StockMovementRequest) error {
	_, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		return err
	}

	newStockMovement := &model.StockMovement{
		ProductID: 	req.ProductID,
		UserID: 	userID,
		Type:      	req.Type,
		Quantity:  	req.Quantity,
		Notes:     	req.Notes,
	}

	return s.stockMovementRepo.Create(newStockMovement)
}

func (s *StockMovementService) FindAllStockMovements() ([]model.StockMovement, error) {
	return s.stockMovementRepo.FindAll()
}

func (s *StockMovementService) FindStockMovementByID(id uint) (*model.StockMovement, error) {
	return s.stockMovementRepo.FindByID(id)
}