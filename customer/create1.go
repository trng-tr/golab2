package customer

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func createCustomer() (Customer, error) {
	var c Customer = Customer{} // init empty customer
	var uuid string = uuid.New().String()
	c.uuid = uuid
	fmt.Print("Saisir le nom du client: ")

	lastname, err := reader.ReadString('\n') // read until back line
	if err != nil {
		return Customer{}, errors.New(errorMgs)
	}
	c.lastname = lastname
	fmt.Print("Saisir le prénom du client: ")
	firstname, err := reader.ReadString('\n')
	if err != nil {
		return Customer{}, errors.New(errorMgs)
	}
	c.firstname = firstname

	fmt.Print("Saisir le email du client: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return Customer{}, errors.New(errorMgs)
	}
	c.email = email
	return c, nil
}
