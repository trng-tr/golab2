package models

import (
	"time"
)

type Customer struct {
	Uuid      string
	Firstname string
	Lastname  string
	Email     string
	Address   Address // remplacé par AddressId référençant l'entité Address
}
type Address struct {
	Uuid       string
	StreetNum  int64
	StreetName string
	ZipCode    string
	City       string
	Country    string
}
type Product struct { // utilsé pour remplir le stock
	Uuid        string
	Sku         string // référence du produit
	Name        string
	Description string
	UnitPrice   float64 // le prix peut changer avec le temps
	Currency    string
	Stock       Stock     // qté du produit en stock
	CreatedAt   time.Time // date and time
	UpdatedAt   time.Time // date and time
	IsActive    bool
}
type Stock struct {
	Quantity int
}
type Order struct {
	Uuid      string
	Customer  Customer
	Items     []OrderItem
	OrderDate string // équivalent en LocalDateTime en Java
}
type OrderItem struct { //utilisé comme dto pour la commande
	ProductName string
	Quantity    int
	UnitPrice   float64 // le prix au momenet de la commande ne change pas
}

const (
	Euro string = "€"
)
