package models

import (
	"errors"
)

func NewOder(id string, c Customer, items []OrderItem, crAt string) (Order, error) {
	if len(items) <= 0 {
		return Order{}, errors.New("order error: commande vide")
	}
	if c.Uuid == "" || c.Firstname == "" || c.Lastname == "" || c.Email == "" {
		return Order{}, errors.New("order error: invalid customer info")
	}

	return Order{
		Uuid:      id,
		Customer:  c,
		Items:     items,
		OrderDate: crAt,
	}, nil
}

func NewOrderItem(productName string, qte int, pce float64) (OrderItem, error) {
	return OrderItem{
		ProductName: productName,
		Quantity:    qte,
		Price:       pce,
	}, nil
}
