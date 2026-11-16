package dto

type SupplierRequest struct {
	Name	string `json:"name" binding:"required"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}