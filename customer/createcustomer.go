package customer

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/trng-tr/golab2/objects"
	"github.com/trng-tr/golab2/scan"
)

func createCustomer() (objects.Customer, error) {
	var c objects.Customer = objects.Customer{} // init empty customer
	var uuid string = uuid.New().String()
	c.uuid = uuid
	fmt.Print("Saisir le nom du client: ")
	if !scan.Scanner.Scan() {
		return objects.Customer{}, errors.New(readingError)
	}
	var lastname string = scan.Scanner.Text()
	if lastname == "" {
		return objects.Customer{}, errors.New(emptyError)
	}
	c.lastname = lastname
	fmt.Print("Saisir le prénom du client: ")
	if !scan.Scanner.Scan() {
		return objects.Customer{}, errors.New(readingError)
	}
	var firstname string = scan.Scanner.Text()
	if firstname == "" {
		return objects.Customer{}, errors.New(emptyError)
	}
	c.firstname = firstname

	fmt.Print("Saisir le email du client: ")
	if !scan.Scanner.Scan() {
		return objects.Customer{}, errors.New(readingError)
	}
	var email string = scan.Scanner.Text()
	if email == "" {
		return objects.Customer{}, errors.New(emptyError)
	}
	c.email = email
	return c, nil
}
