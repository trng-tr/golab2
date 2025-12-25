package services

import (
	"errors"
	"time"

	"github.com/trng-tr/golab2/customer2/models"
)

func CreateOrder(nbItems int) (models.Order, error) {
	if nbItems <= 0 {
		return models.Order{}, errors.New("nompre de produits pour cette cmd invalid")
	}

	var uuid = generateUuid("Order")
	customer, err := CreateCustomer()
	if err != nil {
		return models.Order{}, err
	}

	var items = make([]models.OrderItem, 0, nbItems)
	for i := 1; i <= nbItems; i++ {
		item, err := CreateOrderItem(i)
		if err != nil {
			return models.Order{}, err
		}
		items = append(items, item)
	}
	var createdAt = time.Now().Format("2006-01-02 15:04:05") // formatage avec date de référence
	return models.NewOder(uuid, customer, items, createdAt)
}
