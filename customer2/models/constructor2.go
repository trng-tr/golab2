package models

import (
	"errors"
	"time"
)

func NewProduct(id string, sku string, name string, desc string, pce float64, cur string, stk Stock, crAt time.Time, upAt time.Time, act bool) (Product, error) {
	if len(desc) < 10 {
		return Product{}, errors.New("product error: description incomplete")
	}
	if len(name) < 2 {
		return Product{}, errors.New("product error: le nom du produit trop court")
	}
	if pce < 1 {
		return Product{}, errors.New("product error: le prix invalide")
	}
	if stk.Quantity <= 1 {
		return Product{}, errors.New("product error: qte du produit invalide")
	}

	var p Product = Product{
		Uuid:        id,
		Sku:         sku,
		Name:        name,
		Description: desc,
		Price:       pce,
		Currency:    cur,
		Stock:       stk,
		CreatedAt:   crAt,
		UpdatedAt:   upAt,
		IsActive:    act,
	}

	return p, nil
}
