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

	var uuid string = generateUuid("Order")
	customer, err := CreateCustomer()
	if err != nil {
		return models.Order{}, err
	}

	var items []models.OrderItem = make([]models.OrderItem, 0, nbItems)
	for i := 1; i <= nbItems; i++ {
		item, err := createOrderItem(i)
		if err != nil {
			return models.Order{}, err
		}
		items = append(items, item)
	}
	var createdAt string = time.Now().Format("2026-01-01") // formatage avec date de référence
	return models.NewOder(uuid, customer, items, createdAt)
}
