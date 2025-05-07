package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/GoodsChain/backend/usecase"
	"github.com/GoodsChain/backend/model"
	"net/http"
	"github.com/google/uuid"
)

type SupplierHandler struct {
	supplierUsecase usecase.SupplierUsecase
}

func NewSupplierHandler(supplierUsecase usecase.SupplierUsecase) *SupplierHandler {
	return &SupplierHandler{
		supplierUsecase: supplierUsecase,
	}
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID if not provided
	if supplier.ID == "" {
		supplier.ID = uuid.New().String()
	}

	if err := h.supplierUsecase.CreateSupplier(&supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.supplierUsecase.GetSupplier(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.supplierUsecase.UpdateSupplier(id, &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated successfully"})
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if err := h.supplierUsecase.DeleteSupplier(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}

func (h *SupplierHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.supplierUsecase.GetAllSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suppliers"})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}
