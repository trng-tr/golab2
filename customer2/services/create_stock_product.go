package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/customer2/models"
	"github.com/trng-tr/golab2/read"
)

func createStockProduct(numero int) (models.Product, error) {
	fmt.Printf("saisir le nom du produit %d: ", numero)
	name, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Product{}, errors.New(constantes.ReadingError)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return models.Product{}, errors.New(constantes.EmptyError)
	} else if len(name) < 2 {
		return models.Product{}, errors.New("product error: le nom du produit trop court")
	}
	var uuid = generateUuid(name)
	var sku = generateUuid(name)

	fmt.Printf("saisir la description du produit %d:", numero)
	desc, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Product{}, errors.New(constantes.ReadingError)
	}
	desc = strings.TrimSpace(desc)
	if desc == "" {
		return models.Product{}, errors.New(constantes.EmptyError)
	} else if len(desc) < 1 {
		return models.Product{}, errors.New("product error: description incomplete")
	}
	fmt.Printf("saisir le prix pour le produit %d: ", numero)
	strPrice, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Product{}, errors.New(constantes.ReadingError)
	}
	strPrice = strings.TrimSpace(strPrice)
	if strPrice == "" {
		return models.Product{}, errors.New(constantes.EmptyError)
	}
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		return models.Product{}, errors.New(constantes.ConversionError)
	} else if price <= 0 {
		return models.Product{}, errors.New("product error: le prix invalide")
	}
	fmt.Printf("saisir la quantité pour le produit%d: ", numero)
	str, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Product{}, errors.New(constantes.ReadingError)
	}
	str = strings.TrimSpace(str)
	if str == "" {
		return models.Product{}, errors.New(constantes.EmptyError)
	}

	qte, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return models.Product{}, errors.New(constantes.ConversionError)
	} else if qte <= 0 {
		return models.Product{}, errors.New("product error: qte du produit invalide")
	}
	var stock models.Stock = models.Stock{
		Quantity: int(qte),
	}

	product, err := models.NewProduct(uuid, sku, name, desc, price, models.Euro,
		stock, time.Now(), time.Time{}, true)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func FillStock(nbProducts int) ([]models.Product, error) {
	if nbProducts <= 0 {
		return nil, errors.New("nombre de produits pour le stock est invalid")
	}
	var products []models.Product = make([]models.Product, 0, nbProducts)
	for i := 1; i <= nbProducts; i++ {
		product, err := createStockProduct(i)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
