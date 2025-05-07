package model

import (
	"time"
)

// Car represents a car in the system.
type Car struct {
	ID         string    `json:"id" db:"id" example:"car_01H8ZJ5XQ8X5X8X5X8X5X8X5X8" description:"Unique identifier for the car"`
	Name       string    `json:"name" db:"name" binding:"required" example:"Toyota Camry" description:"Name of the car model"`
	SupplierID string    `json:"supplier_id" db:"supp_id" binding:"required" example:"supp_01H7ZD00X8X5X8X5X8X5X8X5X8" description:"Identifier of the supplier"`
	Price      int       `json:"price" db:"price" binding:"required,gt=0" example:"25000" description:"Price of the car in the smallest currency unit (e.g., cents)"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" example:"2023-03-20T10:00:00Z" format:"date-time" description:"Timestamp of when the car record was created"`
	CreatedBy  string    `json:"created_by" db:"created_by" example:"admin_user" description:"Identifier of the user/process that created the car record"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at" example:"2023-03-21T11:30:00Z" format:"date-time" description:"Timestamp of when the car record was last updated"`
	UpdatedBy  string    `json:"updated_by" db:"updated_by" example:"admin_user" description:"Identifier of the user/process that last updated the car record"`
}
