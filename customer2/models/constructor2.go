package models

import (
	"time"
)

func NewProduct(id string, sku string, name string, desc string, pce float64,
	cur string, stk Stock, crAt time.Time, upAt time.Time, act bool) (Product, error) {

	var p Product = Product{
		Uuid:        id,
		Sku:         sku,
		Name:        name,
		Description: desc,
		UnitPrice:   pce,
		Currency:    cur,
		Stock:       stk,
		CreatedAt:   crAt,
		UpdatedAt:   upAt,
		IsActive:    act,
	}

	return p, nil
}
