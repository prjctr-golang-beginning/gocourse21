package model

import (
	"github.com/google/uuid"
	"time"
)

func NewBrand() *Brand {
	res := &Brand{}

	return res
}

type Brand struct {
	ID        uuid.UUID  `db:"id" table_name:"brands"`
	Name      string     `db:"name"`
	Code      string     `db:"code"`
	IsMain    bool       `db:"is_main"`
	Alias     string     `db:"alias"`
	Order     int        `db:"order"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
