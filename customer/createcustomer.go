package customer

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/scan"
)

func createCustomer() (Customer, error) {
	var c Customer = Customer{} // init empty customer
	var uuid string = uuid.New().String()
	c.uuid = uuid
	fmt.Print("Saisir le nom du client: ")
	if !scan.Scanner.Scan() {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	var lastname string = scan.Scanner.Text()
	if lastname == "" {
		return Customer{}, errors.New(constantes.EmptyError)
	}
	c.lastname = lastname
	fmt.Print("Saisir le prénom du client: ")
	if !scan.Scanner.Scan() {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	var firstname string = scan.Scanner.Text()
	if firstname == "" {
		return Customer{}, errors.New(constantes.EmptyError)
	}
	c.firstname = firstname

	fmt.Print("Saisir le email du client: ")
	if !scan.Scanner.Scan() {
		return Customer{}, errors.New(constantes.ReadingError)
	}
	var email string = scan.Scanner.Text()
	if email == "" {
		return Customer{}, errors.New(constantes.EmptyError)
	}
	c.email = email
	return c, nil
}
