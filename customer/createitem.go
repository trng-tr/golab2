package customer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/scan"
)

func createItem(numItem int) (Item, error) {
	var i = Item{}
	var u = uuid.New().String()
	i.uuid = u
	fmt.Printf("Saisir le titre de l'item %d :", numItem)
	if !scan.Scanner.Scan() {
		return Item{}, errors.New(constantes.ReadingError)
	}
	var title = scan.Scanner.Text()

	if title == "" {
		return Item{}, errors.New(constantes.EmptyError)
	}
	i.title = strings.TrimSpace(title) // remove space au débit et à la fin
	fmt.Printf("Saisir la quantité de l item %d :", numItem)
	if !scan.Scanner.Scan() {
		return Item{}, errors.New(constantes.ReadingError)
	}
	var str = scan.Scanner.Text()
	if str == "" {
		return Item{}, errors.New(constantes.EmptyError)
	}

	quantity, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Item{}, errors.New("error: conversion impossible")
	}
	i.quantity = uint(quantity)
	i.isFood = true
	return i, nil
}
