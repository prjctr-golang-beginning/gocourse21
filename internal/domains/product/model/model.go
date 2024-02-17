package model

import (
	"github.com/google/uuid"
	"time"
)

func NewProduct() *Product {
	res := &Product{}

	return res
}

type Product struct {
	ID        uuid.UUID  `db:"id" table_name:"products"`
	BrandID   uuid.UUID  `db:"brand_id"`
	Status    string     `db:"status"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
