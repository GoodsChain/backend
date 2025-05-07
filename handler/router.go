package handler

import "github.com/gin-gonic/gin"

// InitRoutes sets up all the routes
func InitRoutes(engine *gin.Engine, customerHandler *CustomerHandler, supplierHandler *SupplierHandler, carHandler *CarHandler) {
	// Register global error handling middleware
	engine.Use(ErrorHandlingMiddleware())

	customerGroup := engine.Group("/customers")
	{
		customerGroup.POST("/", customerHandler.CreateCustomer)
		customerGroup.GET("/", customerHandler.GetAllCustomers)
		customerGroup.GET("/:id", customerHandler.GetCustomer)
		customerGroup.PUT("/:id", customerHandler.UpdateCustomer)
		customerGroup.DELETE("/:id", customerHandler.DeleteCustomer)
	}

	supplierGroup := engine.Group("/suppliers")
	{
		supplierGroup.POST("/", supplierHandler.CreateSupplier)
		supplierGroup.GET("/", supplierHandler.GetAllSuppliers)
		supplierGroup.GET("/:id", supplierHandler.GetSupplier)
		supplierGroup.PUT("/:id", supplierHandler.UpdateSupplier)
		supplierGroup.DELETE("/:id", supplierHandler.DeleteSupplier)
	}

	carGroup := engine.Group("/cars")
	{
		carGroup.POST("/", carHandler.CreateCar)
		carGroup.GET("/", carHandler.GetAllCars)
		carGroup.GET("/:id", carHandler.GetCar)
		carGroup.PUT("/:id", carHandler.UpdateCar)
		carGroup.DELETE("/:id", carHandler.DeleteCar)
	}
}
