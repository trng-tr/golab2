package customer

import (
	"fmt"
	"strconv"

	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/scan"
)

func GetOrder() {
	fmt.Print("Saisir le nombre d'items pour ta commande: ")
	if !scan.Scanner.Scan() {
		fmt.Println(constantes.ReadingError)
	}
	var str string = scan.Scanner.Text()
	if str == "" {
		fmt.Println(constantes.EmptyError)
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
