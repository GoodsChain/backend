package model

import (
	"time"
)

type Supplier struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	Address   string    `db:"address" json:"address" binding:"required"`
	Phone     string    `db:"phone" json:"phone"` // Optional
	Email     string    `db:"email" json:"email" binding:"required,email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
