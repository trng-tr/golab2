package models

import "fmt"

func (o *Order) PrintOrder() {
	fmt.Println("{")
	fmt.Printf("  id: %s\n", o.Uuid)
	fmt.Println("  customer {")
	fmt.Printf("     id: %s\n", o.Customer.Uuid)
	fmt.Printf("     prénom: %s\n", o.Customer.Firstname)
	fmt.Printf("     nom: %s\n", o.Customer.Lastname)
	fmt.Printf("     email: %s\n", o.Customer.Email)
	fmt.Println("    address {")
	fmt.Printf("       id: %s\n", o.Customer.Address.Uuid)
	fmt.Printf("       numero: %d\n", o.Customer.Address.StreetNum)
	fmt.Printf("       nom de la rue: %s\n", o.Customer.Address.StreetName)
	fmt.Printf("       code postal: %s\n", o.Customer.Address.ZipCode)
	fmt.Printf("       ville: %s\n", o.Customer.Address.City)
	fmt.Printf("       pays: %s\n", o.Customer.Address.Country)
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println("  items: [")
	for i, item := range o.Items {
		fmt.Printf("   item:%d : libelé: %s, quantité: %d, pu: %.2f, pt: %.2f\n",
			i+1,
			item.ProductName,
			item.Quantity,
			item.UnitPrice,
			item.UnitPrice*float64(item.Quantity),
		)
	}
	fmt.Println("  ]")

	fmt.Printf("  orderDate: %s\n", o.OrderDate)
	fmt.Println("}")
}

// AddItem outils mis en disposition de l'utilisateur pour ajouter des items dans sa commande
func (o *Order) AddItem(orderItem OrderItem) {
	o.Items = append(o.Items, orderItem)
}
