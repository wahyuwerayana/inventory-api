package dto

import "inventory-api/internal/model"

type StockMovementRequest struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=IN OUT"` // "IN" or "OUT"
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
	Notes     string `json:"notes"`
}

type StockMovementResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	ProductSKU  string `json:"product_sku"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
	RecordedBy  string `json:"recorded_by"`
	Notes       string `json:"notes"`
}

func ToStockMovementResponse(s *model.StockMovement) *StockMovementResponse {
	return &StockMovementResponse{
		ID:          s.ID,
		ProductName: s.Product.Name,
		ProductSKU:  s.Product.SKU,
		Type:        s.Type,
		Quantity:    s.Quantity,
		RecordedBy:  s.User.Name,
		Notes:       s.Notes,
	}
}