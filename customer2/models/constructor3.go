package models

func NewOder(id string, c Customer, items []OrderItem, crAt string) (Order, error) {

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
		UnitPrice:   pce,
	}, nil
}
