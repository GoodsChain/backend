package handler

import "github.com/gin-gonic/gin"

// InitRoutes sets up all the routes
// Now accepts a RouterGroup instead of Engine to support API versioning
func InitRoutes(router gin.IRouter, customerHandler *CustomerHandler, supplierHandler *SupplierHandler,
	carHandler *CarHandler, customerCarHandler *CustomerCarHandler) {
	// Note: global middleware should be registered at the engine level, not here

	customerGroup := router.Group("/customers")
	{
		customerGroup.POST("", customerHandler.CreateCustomer)
		customerGroup.GET("", customerHandler.GetAllCustomers)
		customerGroup.GET("/:id", customerHandler.GetCustomer)
		customerGroup.PUT("/:id", customerHandler.UpdateCustomer)
		customerGroup.DELETE("/:id", customerHandler.DeleteCustomer)
		customerGroup.GET("/:customer_id/cars", customerCarHandler.GetByCustomerID)
	}

	supplierGroup := router.Group("/suppliers")
	{
		supplierGroup.POST("", supplierHandler.CreateSupplier)
		supplierGroup.GET("", supplierHandler.GetAllSuppliers)
		supplierGroup.GET("/:id", supplierHandler.GetSupplier)
		supplierGroup.PUT("/:id", supplierHandler.UpdateSupplier)
		supplierGroup.DELETE("/:id", supplierHandler.DeleteSupplier)
	}

	carGroup := router.Group("/cars")
	{
		carGroup.POST("", carHandler.CreateCar)
		carGroup.GET("", carHandler.GetAllCars)
		carGroup.GET("/:id", carHandler.GetCar)
		carGroup.PUT("/:id", carHandler.UpdateCar)
		carGroup.DELETE("/:id", carHandler.DeleteCar)
		carGroup.GET("/:car_id/customers", customerCarHandler.GetByCarID)
	}

	customerCarGroup := router.Group("/customer-cars")
	{
		customerCarGroup.POST("", customerCarHandler.Create)
		customerCarGroup.GET("", customerCarHandler.GetAll)
		customerCarGroup.GET("/:id", customerCarHandler.GetByID)
		customerCarGroup.PUT("/:id", customerCarHandler.Update)
		customerCarGroup.DELETE("/:id", customerCarHandler.Delete)
	}
}
