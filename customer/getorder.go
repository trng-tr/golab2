package customer

import (
	"fmt"
	"strconv"
)

func GetOrder() {
	fmt.Print("Saisir le nombre d'items pour ta commande: ")
	if !scanner.Scan() {
		fmt.Println(readingError)
	}
	var str string = scanner.Text()
	if str == "" {
		fmt.Println(emptyError)
		return
	}
	nbItems, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println("error conversion")
	}
	order, err := createOrder(int(nbItems))
	if err != nil {
		fmt.Println("error is raised: ", err)
		return
	}
	fmt.Println("la commande:", order)
}
