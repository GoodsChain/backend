package handler

import (
	"net/http"

	"github.com/GoodsChain/backend/model"
	"github.com/GoodsChain/backend/usecase"
	"github.com/gin-gonic/gin"
)

// CustomerCarHandler manages HTTP requests related to customer-car relationships
type CustomerCarHandler struct {
	CustomerCarUsecase usecase.CustomerCarUsecase
}

// NewCustomerCarHandler creates a new instance of CustomerCarHandler
func NewCustomerCarHandler(customerCarUsecase usecase.CustomerCarUsecase) *CustomerCarHandler {
	return &CustomerCarHandler{
		CustomerCarUsecase: customerCarUsecase,
	}
}

// Create godoc
// @Summary Create customer car relationship
// @Description Create a new customer car relationship
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param customerCar body model.CustomerCar true "Customer car data"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customer-cars [post]
func (h *CustomerCarHandler) Create(c *gin.Context) {
	var customerCar model.CustomerCar
	if err := c.ShouldBindJSON(&customerCar); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.CustomerCarUsecase.CreateCustomerCar(&customerCar); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customerCar)
}

// GetByID godoc
// @Summary Get customer car by ID
// @Description Get a customer car relationship by its ID
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param id path string true "Customer Car ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customer-cars/{id} [get]
func (h *CustomerCarHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	customerCar, err := h.CustomerCarUsecase.GetCustomerCar(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Customer car relationship not found"})
		return
	}

	c.JSON(http.StatusOK, customerCar)
}

// GetAll godoc
// @Summary Get all customer car relationships
// @Description Get all customer car relationships
// @Tags customer-cars
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customer-cars [get]
func (h *CustomerCarHandler) GetAll(c *gin.Context) {
	customerCars, err := h.CustomerCarUsecase.GetAllCustomerCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to get customer car relationships"})
		return
	}

	c.JSON(http.StatusOK, customerCars)
}

// GetByCustomerID godoc
// @Summary Get customer cars by customer ID
// @Description Get all cars owned by a specific customer
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customers/{customer_id}/cars [get]
func (h *CustomerCarHandler) GetByCustomerID(c *gin.Context) {
	customerID := c.Param("customer_id")

	customerCars, err := h.CustomerCarUsecase.GetCustomerCarsByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to get customer car relationships"})
		return
	}

	c.JSON(http.StatusOK, customerCars)
}

// GetByCarID godoc
// @Summary Get customer cars by car ID
// @Description Get all customers who own a specific car
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param car_id path string true "Car ID"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/cars/{car_id}/customers [get]
func (h *CustomerCarHandler) GetByCarID(c *gin.Context) {
	carID := c.Param("car_id")

	customerCars, err := h.CustomerCarUsecase.GetCustomerCarsByCarID(carID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to get customer car relationships"})
		return
	}

	c.JSON(http.StatusOK, customerCars)
}

// Update godoc
// @Summary Update customer car
// @Description Update a customer car relationship
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param id path string true "Customer Car ID"
// @Param customerCar body model.CustomerCar true "Customer car data"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customer-cars/{id} [put]
func (h *CustomerCarHandler) Update(c *gin.Context) {
	id := c.Param("id")
	
	var customerCar model.CustomerCar
	if err := c.ShouldBindJSON(&customerCar); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.CustomerCarUsecase.UpdateCustomerCar(id, &customerCar); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Customer car relationship updated successfully"})
}

// Delete godoc
// @Summary Delete customer car
// @Description Delete a customer car relationship
// @Tags customer-cars
// @Accept json
// @Produce json
// @Param id path string true "Customer Car ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/customer-cars/{id} [delete]
func (h *CustomerCarHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.CustomerCarUsecase.DeleteCustomerCar(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Customer car relationship deleted successfully"})
}
