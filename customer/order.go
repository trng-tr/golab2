package customer

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func createOrder(nbItems int) (Order, error) {
	var order Order = Order{}
	order.uuid = uuid.New().String() // generer uuid from github package
	if nbItems <= 0 {
		return Order{}, errors.New("Nombre d'item invalid")
	}

	var items []Item = make([]Item, 0, nbItems)

	for i := 0; i < nbItems; i++ {
		item, err := createItem()
		if err != nil {
			return Order{}, err
		}
		addItem(item, items)
	}
	var client Customer = createCustomer()
	order.customer = client

	return order, nil

}

func addItem(item Item, items []Item) {
	items = append(items, item)
}

func GetOrder() {
	var nbItems int
	fmt.Print("Saisir le nombre d'items pour ta commande: ")
	fmt.Scan(&nbItems)
	order, err := createOrder(nbItems)
	if err != nil {
		fmt.Println("error is raised: ", err)
		return
	}
	fmt.Println("la commande:", order)
}
