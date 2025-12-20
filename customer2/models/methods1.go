package models

import (
	"fmt"
)

func (c Customer) PrintCustomer() { // methode pour affichir un customer
	fmt.Println("{")
	fmt.Println(" customer {")
	fmt.Printf("     id: %s\n", c.Uuid)
	fmt.Printf("     prénom: %s\n", c.Firstname)
	fmt.Printf("     nom: %s\n", c.Lastname)
	fmt.Printf("     email: %s\n", c.Email)
	fmt.Println("    address {")
	fmt.Printf("       id: %s\n", c.Address.Uuid)
	fmt.Printf("       numero: %d\n", c.Address.StreetNum)
	fmt.Printf("       nom de la rue: %s\n", c.Address.StreetName)
	fmt.Printf("       code postal: %s\n", c.Address.ZipCode)
	fmt.Printf("       ville: %s\n", c.Address.City)
	fmt.Printf("       pays: %s\n", c.Address.Country)
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println("}")
}

func (pdt Product) PrintProduct() { // methode pour affichir un product
	var updated string
	if pdt.UpdatedAt.IsZero() {
		updated = "null"
	} else {
		updated = pdt.UpdatedAt.Format("02-01-2006 15:04")
	}
	fmt.Printf(
		"{\n"+
			"  id: %s,\n"+
			"  sku: %s,\n"+
			"  name: %s,\n"+
			"  desc: %s,\n"+
			"  pu: %.2f %s,\n"+
			"  pt: %.2f %s,\n"+
			"  stock: %d,\n"+
			"  created at: %s,\n"+
			"  updated at: %s,\n"+
			"  is_active: %v\n"+
			"}",
		pdt.Uuid, pdt.Sku, pdt.Name, pdt.Description, pdt.UnitPrice, pdt.Currency,
		pdt.UnitPrice*float64(pdt.Stock.Quantity), pdt.Currency, pdt.Stock.Quantity,
		pdt.CreatedAt.Format("2006-01-02 15:04:05"), updated, pdt.IsActive,
	)
}
