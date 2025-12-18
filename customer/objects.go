package customer

type Customer struct {
	uuid      string
	firstname string
	lastname  string
	email     string
}

type Item struct {
	uuid     string
	title    string
	quantity uint
	isFood   bool
}

type Order struct {
	uuid     string
	items    []Item
	customer Customer
}
