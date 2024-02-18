package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"gocourse21/internal/core/db"
	"gocourse21/internal/domains/brand/model"
	model2 "gocourse21/internal/domains/product/model"
	model3 "gocourse21/internal/domains/supplier/model"
	"gocourse21/internal/lib/generator"
)

func main() {
	b := model.NewBrand()

	spew.Dump(b)

	db.PopulateWith(b, map[string]any{
		`id`:         uuid.New(),
		`is_main`:    `true`,
		`name`:       `Test Brand`,
		`code`:       `Test Brand`,
		`alias`:      `test_brand`,
		`order`:      27,
		`created_at`: `2024-02-17 17:49:52`,
	})

	spew.Dump(b)

	generator.GenerateServiceFiles(model.NewBrand(), model2.NewProduct(), model3.NewSupplier())
	generator.GenerateRepositoryFiles(model.NewBrand(), model2.NewProduct(), model3.NewSupplier())
}
