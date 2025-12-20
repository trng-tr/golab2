package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trng-tr/golab2/constantes"
	"github.com/trng-tr/golab2/customer2/services"
	"github.com/trng-tr/golab2/read"
)

func main() {
	/* customer.GetOrder()
	c, err := services.CreateCustomer()
	if err != nil {
		fmt.Println("error is raised: ", err)
		return
	}
	c.PrintCustomer() */
	fmt.Println("On va remplir le stock avec quelques produits")
	fmt.Print("Saisir le nombre de produits pour le stock :")
	str, err := read.StreamReader.ReadString('\n') // read until meet \n
	if err != nil {
		fmt.Println(constantes.ReadingError)
		return
	}
	str = strings.TrimSpace(str)
	if str == "" {
		fmt.Println(constantes.EmptyError)
	}
	nbProducts, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(constantes.ConversionError)
	}
	stock, err := services.FillStcok(int(nbProducts))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, product := range stock {
		product.PrintProduct()
	}
	fmt.Print("On passe une commande, saisir le nombre de produits:")
	strNbProd, err := read.StreamReader.ReadString('\n') // read until meet \n
	if err != nil {
		fmt.Println(constantes.ReadingError)
		return
	}
	strNbProd = strings.TrimSpace(strNbProd)
	if strNbProd == "" {
		fmt.Println(constantes.EmptyError)
		return
	}
	cmdNbProducts, err := strconv.Atoi(strNbProd)
	if err != nil {
		fmt.Println(constantes.ConversionError)
		return
	}
	order, err := services.CreateOrder(int(cmdNbProducts))
	if err != nil {
		fmt.Println(err)
		return
	}
	order.PrintOrder()
	item, err := services.CreateOrderItem(len(order.Items))
	if err != nil {
		fmt.Println(err)
		return
	}
	order.AddItem(item)
	order.PrintOrder()
}
