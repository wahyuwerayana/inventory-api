package handler

import (
	"fmt"
	"inventory-api/internal/dto"
	"inventory-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	supplierService *service.SupplierService
}

func NewSupplierHandler(s *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{supplierService: s}
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var req dto.SupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.supplierService.CreateSupplier(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Supplier created successfully",
		"supplier": req,
	})
}

func (h *SupplierHandler) FindAllSuppliers(c *gin.Context) {
	suppliers, err := h.supplierService.FindAllSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"suppliers": suppliers,
	})
}

func (h *SupplierHandler) FindSupplierByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid supplier ID",
		})
		return
	}

	supplier, err := h.supplierService.FindSupplierByID((uint)(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplier,
	})
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid supplier ID",
		})
		return
	}

	var req dto.SupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.supplierService.UpdateSupplier((uint)(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Supplier %d updated successfully", id),
	})
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid supplier ID",
		})
		return
	}

	err = h.supplierService.DeleteSupplier((uint)(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Supplier %d deleted successfully", id),
	})
}