package handler

import (
	"inventory-api/internal/dto"
	"inventory-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockMovementHandler struct {
	stockMovementService *service.StockMovementService
}

func NewStockMovementHandler(s *service.StockMovementService) *StockMovementHandler {
	return &StockMovementHandler{stockMovementService: s}
}

func (h *StockMovementHandler) CreateStockMovement(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in context"})
		return
	}

	var req *dto.StockMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.stockMovementService.CreateStockMovement(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Stock movement created successfully"})
}

func (h *StockMovementHandler) FindAllStockMovements(c *gin.Context) {
	stockMovements, err := h.stockMovementService.FindAllStockMovements()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stockMovementResponses := make([]dto.StockMovementResponse, len(stockMovements))
	for i, sm := range stockMovements {
		stockMovementResponses[i] = *dto.ToStockMovementResponse(&sm)
	}

	c.JSON(http.StatusOK, gin.H{
		"stockMovements": stockMovementResponses,
	})
}

func (h *StockMovementHandler) FindStockMovementByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	stockMovement, err := h.stockMovementService.FindStockMovementByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stockMovement": dto.ToStockMovementResponse(stockMovement),
	})
}