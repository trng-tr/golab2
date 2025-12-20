package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/customer2/models"
	"github.com/trng-tr/golab2/read"
)

func CreateCustomer() (models.Customer, error) {
	fmt.Print("Saisir le nom du client: ")
	lastname, err := read.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return models.Customer{}, errors.New(constantes.ReadingError)
	}
	lastname = strings.TrimSpace(lastname)
	if lastname == "" {
		return models.Customer{}, errors.New(constantes.EmptyError)
	}
	var uuid string = generateUuid(lastname)

	fmt.Print("Saisir le prénom du client: ")
	firstname, err := read.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return models.Customer{}, errors.New(constantes.ReadingError)
	}
	firstname = strings.TrimSpace(firstname)
	if firstname == "" {
		return models.Customer{}, errors.New(constantes.EmptyError)
	}
	fmt.Print("Saisir le email  du client: ")
	email, err := read.StreamReader.ReadString('\n') //read until meet \n
	if err != nil {
		return models.Customer{}, errors.New(constantes.ReadingError)
	}
	email = strings.TrimSpace(email)
	if email == "" {
		return models.Customer{}, errors.New(constantes.EmptyError)
	} else if len(email) < 5 || !strings.Contains(email, "@") {
		return models.Customer{}, errors.New("Email invalid")
	}
	address, err := createAddress()
	if err != nil {
		return models.Customer{}, err
	}
	c, err := models.NewCustomer(uuid, firstname, lastname, email, address)
	if err != nil {
		return models.Customer{}, err
	}
	return c, nil
}
