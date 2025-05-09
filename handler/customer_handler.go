package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/GoodsChain/backend/usecase"
	"github.com/GoodsChain/backend/model"
	"net/http"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Adds a new customer to the system. The ID is auto-generated if not provided.
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer body model.Customer true "Customer object to be created"
// @Success 201 {object} model.Customer "Successfully created customer"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer
	// Note: For request, ID, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy are typically ignored or server-set.
	// The model.Customer is used here for simplicity; a dedicated CreateCustomerRequest struct could be used.
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: "invalid_request", Message: err.Error()})
		return
	}

	// Generate UUID if not provided
	if customer.ID == "" {
		customer.ID = uuid.New().String()
	}

	if err := h.customerUsecase.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// GetCustomer godoc
// @Summary Get a customer by ID
// @Description Retrieves a customer's details based on their unique ID.
// @Tags Customers
// @Produce json
// @Param id path string true "Customer ID" example:"cust_01H7ZCN4X8X5X8X5X8X5X8X5X8"
// @Success 200 {object} model.Customer "Successfully retrieved customer"
// @Failure 404 {object} model.ErrorResponse "Customer not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.customerUsecase.GetCustomer(id)
	if err != nil {
		// Assuming GetCustomer returns a specific error type that can be checked for "not found"
		// For now, using the existing logic which might be improved in usecase layer.
		c.JSON(http.StatusNotFound, model.ErrorResponse{Code: "not_found", Message: "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer godoc
// @Summary Update an existing customer
// @Description Updates the details of an existing customer identified by their ID.
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID" example:"cust_01H7ZCN4X8X5X8X5X8X5X8X5X8"
// @Param customer body model.Customer true "Customer object with updated details"
// @Success 200 {object} model.SuccessResponse "Customer updated successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 404 {object} model.ErrorResponse "Customer not found (if ID in body differs or not found by usecase)"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer model.Customer
	// Note: For request, ID, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy are typically ignored or server-set.
	// The model.Customer is used here for simplicity; a dedicated UpdateCustomerRequest struct could be used.
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: "invalid_request", Message: err.Error()})
		return
	}

	// It's good practice to ensure the ID in path matches ID in body if present, or usecase handles it.
	// For now, assuming usecase uses the path `id`.
	if err := h.customerUsecase.UpdateCustomer(id, &customer); err != nil {
		// This could be a not found error or other internal error.
		// Usecase should return distinguishable errors.
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Customer updated successfully"})
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Deletes a customer from the system based on their unique ID.
// @Tags Customers
// @Produce json
// @Param id path string true "Customer ID" example:"cust_01H7ZCN4X8X5X8X5X8X5X8X5X8"
// @Success 200 {object} model.SuccessResponse "Customer deleted successfully"
// @Failure 404 {object} model.ErrorResponse "Customer not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := h.customerUsecase.DeleteCustomer(id); err != nil {
		// Usecase should return distinguishable errors for not found vs internal.
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Customer deleted successfully"})
}

// GetAllCustomers godoc
// @Summary Get all customers
// @Description Retrieves a list of all customers in the system.
// @Tags Customers
// @Produce json
// @Success 200 {array} model.Customer "Successfully retrieved list of customers"
// @Failure 500 {object} model.ErrorResponse "Failed to retrieve customers"
// @Router /customers [get]
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.customerUsecase.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: "internal_error", Message: "Failed to retrieve customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}
