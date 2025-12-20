package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/customer2/models"
	"github.com/trng-tr/golab2/read"
)

func CreateOrderItem(numero int) (models.OrderItem, error) {
	fmt.Printf("saisir le nom du produit  %d à commander: ", numero)
	productName, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.OrderItem{}, errors.New(constantes.ReadingError)
	}
	productName = strings.TrimSpace(productName)
	if productName == "" {
		return models.OrderItem{}, errors.New(constantes.EmptyError)
	}
	fmt.Printf("saisir la quantité pour le produit  %d a commander: ", numero)
	strQte, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.OrderItem{}, errors.New(constantes.ReadingError)
	}
	strQte = strings.TrimSpace(strQte)
	if strQte == "" {
		return models.OrderItem{}, errors.New(constantes.EmptyError)
	}
	qte, err := strconv.ParseInt(strQte, 10, 64)
	if err != nil {
		return models.OrderItem{}, errors.New(constantes.ConversionError)
	}
	fmt.Printf("saisir le prix pour le produit  %d a commander: ", numero)
	strPrice, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.OrderItem{}, errors.New(constantes.ReadingError)
	}
	strPrice = strings.TrimSpace(strPrice)

	if strPrice == "" {
		return models.OrderItem{}, errors.New(constantes.EmptyError)
	}
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		return models.OrderItem{}, errors.New(constantes.ConversionError)
	}
	return models.NewOrderItem(productName, int(qte), price)
}
