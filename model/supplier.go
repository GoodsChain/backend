package model

import (
	"time"
)

// Supplier represents a supplier in the system.
type Supplier struct {
	ID        string    `db:"id" json:"id" example:"supp_01H7ZD00X8X5X8X5X8X5X8X5X8" description:"Unique identifier for the supplier"`
	Name      string    `db:"name" json:"name" binding:"required" example:"Supplier Inc." description:"Name of the supplier"`
	Address   string    `db:"address" json:"address" binding:"required" example:"456 Industrial Rd, Factory City, USA" description:"Address of the supplier"`
	Phone     string    `db:"phone" json:"phone" example:"555-987-6543" description:"Phone number of the supplier (optional)"`
	Email     string    `db:"email" json:"email" binding:"required,email" example:"contact@supplierinc.com" description:"Email address of the supplier"`
	CreatedAt time.Time `db:"created_at" json:"created_at" example:"2023-02-10T09:15:00Z" format:"date-time" description:"Timestamp of when the supplier was created"`
	CreatedBy string    `db:"created_by" json:"created_by" example:"system_user" description:"Identifier of the user/process that created the supplier"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" example:"2023-02-11T14:45:00Z" format:"date-time" description:"Timestamp of when the supplier was last updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by" example:"system_user" description:"Identifier of the user/process that last updated the supplier"`
}
