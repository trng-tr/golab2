package customer

import (
	"errors"

	"github.com/google/uuid"
)

func createOrder(nbItems int) (Order, error) {
	var order Order = Order{}
	order.uuid = uuid.New().String() // generer uuid from github package
	if nbItems <= 0 {
		return Order{}, errors.New("nombre d'item invalid")
	}

	var items []Item = make([]Item, 0, nbItems)

	for i := 0; i < nbItems; i++ {
		item, err := createItem()
		if err != nil {
			return Order{}, err
		}
		addItem(item, items)
	}
	client, err := createCustomer()
	if err != nil {
		return Order{}, err
	}
	order.customer = client

	return order, nil

}

func addItem(item Item, items []Item) {
	items = append(items, item)
}
