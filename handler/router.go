package handler

import "github.com/gin-gonic/gin"

// InitRoutes sets up all the customer routes
func InitRoutes(engine *gin.Engine, customerHandler *CustomerHandler) {
	customerGroup := engine.Group("/customers")
	{
		customerGroup.POST("/", customerHandler.CreateCustomer)
		customerGroup.GET("/", customerHandler.GetAllCustomers)  // New route for getting all customers
		customerGroup.GET("/:id", customerHandler.GetCustomer)
		customerGroup.PUT("/:id", customerHandler.UpdateCustomer)
		customerGroup.DELETE("/:id", customerHandler.DeleteCustomer)
	}
}
