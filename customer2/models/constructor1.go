package models

func NewCustomer(uuid string, firstname string, lastanme string, email string, address Address) (Customer, error) {

	var c Customer = Customer{
		Uuid:      uuid,
		Firstname: firstname,
		Lastname:  lastanme,
		Email:     email,
		Address:   address,
	}

	return c, nil
}

func NewAddress(uuid string, streetNum int64, streetName string, zipCode string, city string, country string) (Address, error) {

	return Address{
		Uuid:       uuid,
		StreetNum:  streetNum,
		StreetName: streetName,
		ZipCode:    zipCode,
		City:       city,
		Country:    country,
	}, nil
}
