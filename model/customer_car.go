package model

import (
	"time"
)

// CustomerCar represents a relationship between a customer and a car in the system.
type CustomerCar struct {
	ID        string    `json:"id" db:"id" example:"cc_01H9ZJ5XQ8X5X8X5X8X5X8X5X8" description:"Unique identifier for the customer-car relationship"`
	CarID     string    `json:"car_id" db:"car_id" binding:"required" example:"car_01H8ZJ5XQ8X5X8X5X8X5X8X5X8" description:"Identifier of the car"`
	CustomerID string   `json:"customer_id" db:"cust_id" binding:"required" example:"cust_01H7ZCN4X8X5X8X5X8X5X8X5X8" description:"Identifier of the customer"`
	CreatedAt time.Time `json:"created_at" db:"created_at" example:"2023-03-20T10:00:00Z" format:"date-time" description:"Timestamp of when the record was created"`
	CreatedBy string    `json:"created_by" db:"created_by" example:"admin_user" description:"Identifier of the user/process that created the record"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" example:"2023-03-21T11:30:00Z" format:"date-time" description:"Timestamp of when the record was last updated"`
	UpdatedBy string    `json:"updated_by" db:"updated_by" example:"admin_user" description:"Identifier of the user/process that last updated the record"`
}
