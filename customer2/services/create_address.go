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

func createAddress() (models.Address, error) {
	fmt.Print("Saisir le numero: ")
	str, err := read.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return models.Address{}, errors.New(constantes.ReadingError)
	}

	str = strings.TrimSpace(str)
	if str == "" {
		return models.Address{}, errors.New(constantes.EmptyError)
	}

	streetNum, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return models.Address{}, errors.New(constantes.ConversionError)
	} else if streetNum < 1 {
		return models.Address{}, errors.New("address error: numero invalide")
	}
	fmt.Print("Saisir le nom de la rue: ")
	streetName, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Address{}, errors.New(constantes.ReadingError)
	}
	streetName = strings.TrimSpace(streetName)
	if streetName == "" {
		return models.Address{}, errors.New(constantes.EmptyError)
	}

	var uuid string = generateUuid(streetName)

	fmt.Print("Saisir la boite postale: ")
	zipCode, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Address{}, errors.New(constantes.ReadingError)
	}

	zipCode = strings.TrimSpace(zipCode)
	if zipCode == "" {
		return models.Address{}, errors.New(constantes.EmptyError)
	}
	fmt.Print("Saisir le nom de la ville: ")
	city, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Address{}, errors.New(constantes.ReadingError)
	}
	city = strings.TrimSpace(city)
	if city == "" {
		return models.Address{}, errors.New(constantes.EmptyError)
	}
	fmt.Print("Saisir le nom du pays: ")
	country, err := read.StreamReader.ReadString('\n')
	if err != nil {
		return models.Address{}, errors.New(constantes.ReadingError)
	}
	country = strings.TrimSpace(country)
	if country == "" {
		return models.Address{}, errors.New(constantes.EmptyError)
	}
	return models.NewAddress(uuid, streetNum, streetName, zipCode, city, country)
}
