package model

import (
	"github.com/google/uuid"
	"time"
)

func NewSupplier() *Supplier {
	return &Supplier{}
}

type Supplier struct {
	ID             uuid.UUID  `json:"id" table_name:"suppliers"`
	Name           string     `json:"name"`
	Alias          string     `json:"alias"`
	IsEnabled      bool       `json:"is_enabled"`
	IsDropshipping bool       `json:"is_dropshipping"`
	AutoRequest    bool       `json:"auto_request"`
	PriceActive    bool       `json:"price_active"`
	City           *string    `json:"city"`
	Address        *string    `json:"address"`
	PostCode       *string    `json:"post_code"`
	Label          string     `json:"label"`
	ParentID       *uuid.UUID `json:"parent_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
