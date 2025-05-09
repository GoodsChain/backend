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

// CreateSupplier godoc
// @Summary Create a new supplier
// @Description Adds a new supplier to the system. The ID is auto-generated if not provided.
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param supplier body model.Supplier true "Supplier object to be created"
// @Success 201 {object} model.Supplier "Successfully created supplier"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /suppliers [post]
func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: "invalid_request", Message: err.Error()})
		return
	}

	// Generate UUID if not provided
	if supplier.ID == "" {
		supplier.ID = uuid.New().String()
	}

	if err := h.supplierUsecase.CreateSupplier(&supplier); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// GetSupplier godoc
// @Summary Get a supplier by ID
// @Description Retrieves a supplier's details based on their unique ID.
// @Tags Suppliers
// @Produce json
// @Param id path string true "Supplier ID" example:"supp_01H7ZD00X8X5X8X5X8X5X8X5X8"
// @Success 200 {object} model.Supplier "Successfully retrieved supplier"
// @Failure 404 {object} model.ErrorResponse "Supplier not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /suppliers/{id} [get]
func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.supplierUsecase.GetSupplier(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Code: "not_found", Message: "Supplier not found"})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// UpdateSupplier godoc
// @Summary Update an existing supplier
// @Description Updates the details of an existing supplier identified by their ID.
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID" example:"supp_01H7ZD00X8X5X8X5X8X5X8X5X8"
// @Param supplier body model.Supplier true "Supplier object with updated details"
// @Success 200 {object} model.SuccessResponse "Supplier updated successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 404 {object} model.ErrorResponse "Supplier not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /suppliers/{id} [put]
func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: "invalid_request", Message: err.Error()})
		return
	}

	if err := h.supplierUsecase.UpdateSupplier(id, &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Supplier updated successfully"})
}

// DeleteSupplier godoc
// @Summary Delete a supplier
// @Description Deletes a supplier from the system based on their unique ID.
// @Tags Suppliers
// @Produce json
// @Param id path string true "Supplier ID" example:"supp_01H7ZD00X8X5X8X5X8X5X8X5X8"
// @Success 200 {object} model.SuccessResponse "Supplier deleted successfully"
// @Failure 404 {object} model.ErrorResponse "Supplier not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /suppliers/{id} [delete]
func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if err := h.supplierUsecase.DeleteSupplier(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Supplier deleted successfully"})
}

// GetAllSuppliers godoc
// @Summary Get all suppliers
// @Description Retrieves a list of all suppliers in the system.
// @Tags Suppliers
// @Produce json
// @Success 200 {array} model.Supplier "Successfully retrieved list of suppliers"
// @Failure 500 {object} model.ErrorResponse "Failed to retrieve suppliers"
// @Router /suppliers [get]
func (h *SupplierHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.supplierUsecase.GetAllSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: "Failed to retrieve suppliers"})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}
