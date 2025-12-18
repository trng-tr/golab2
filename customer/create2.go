package customer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func createItem() (Item, error) {
	var i Item = Item{}
	var uuid string = uuid.New().String()
	i.uuid = uuid
	fmt.Print("Saisir le titre de l item: ")
	title, err := reader.ReadString('\n')

	if err != nil {
		return Item{}, errors.New(errorMgs)
	}
	i.title = strings.TrimSpace(title) // remove space au débit et à la fin
	fmt.Print("Saisir la quantité de l item: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return Item{}, errors.New(errorMgs)
	}

	quantity, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Item{}, errors.New("error: conversion impossible")
	}
	i.quantity = uint(quantity)
	i.isFood = true
	return i, nil
}
