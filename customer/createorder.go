package customer

import (
	"errors"

	"github.com/google/uuid"
)

func createOrder(nbItems int) (Order, error) {
	var order = Order{}
	order.uuid = uuid.New().String() // generer uuid from github package
	if nbItems <= 0 {
		return Order{}, errors.New("nombre d'item invalid")
	}

	var items = make([]Item, 0, nbItems)

	for i := 1; i <= nbItems; i++ {
		item, err := createItem(i)
		if err != nil {
			return Order{}, err
		}
		items = append(items, item)
	}
	order.items = items
	client, err := createCustomer()
	if err != nil {
		return Order{}, err
	}

	order.customer = client

	return order, nil
}
