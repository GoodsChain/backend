package model

import (
	"time"
)

// Customer represents a customer in the system.
type Customer struct {
	ID        string    `db:"id" json:"id" example:"cust_01H7ZCN4X8X5X8X5X8X5X8X5X8" description:"Unique identifier for the customer"`
	Name      string    `db:"name" json:"name" binding:"required" example:"John Doe" description:"Name of the customer"`
	Address   string    `db:"address" json:"address" binding:"required" example:"123 Main St, Anytown, USA" description:"Address of the customer"`
	Phone     string    `db:"phone" json:"phone" example:"555-123-4567" description:"Phone number of the customer (optional)"`
	Email     string    `db:"email" json:"email" binding:"required,email" example:"john.doe@example.com" description:"Email address of the customer"`
	CreatedAt time.Time `db:"created_at" json:"created_at" example:"2023-01-15T10:30:00Z" format:"date-time" description:"Timestamp of when the customer was created"`
	CreatedBy string    `db:"created_by" json:"created_by" example:"system_user" description:"Identifier of the user/process that created the customer"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" example:"2023-01-16T11:00:00Z" format:"date-time" description:"Timestamp of when the customer was last updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by" example:"system_user" description:"Identifier of the user/process that last updated the customer"`
}
