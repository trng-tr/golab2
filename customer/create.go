package customer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func createCustomer() Customer {
	var c Customer = Customer{} // init empty customer
	var uuid string = uuid.New().String()
	c.uuid = uuid
	fmt.Print("Saisir le nom du client: ")
	fmt.Scan(&c.lastname)
	fmt.Print("Saisir le prénom du client: ")
	fmt.Scan(&c.firstname)
	fmt.Print("Saisir le email du client: ")
	fmt.Scan(&c.email)

	return c
}

func createItem() (Item, error) {
	var i Item = Item{}
	var uuid string = uuid.New().String()
	i.uuid = uuid
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Saisir le titre de l item: ")
	foodTitle, err := reader.ReadString('\n') // lire jusqu'à rencontrer un retour à la ligne
	if err != nil {
		return Item{}, errors.New("erreur de lecture")
	}
	i.title = strings.TrimSpace(foodTitle) // remove space au débit et à la fin
	fmt.Print("Saisir la qauntité de l item: ")
	fmt.Scan(&i.quantity)
	i.isFood = true
	return i, nil

}
