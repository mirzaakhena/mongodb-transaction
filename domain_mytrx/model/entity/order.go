package entity

import "fmt"

type Order struct {
	ProductName string
	Quantity    int
}

func NewOrder(name string, qty int) (*Order, error) {

	if qty <= 0 {
		return nil, fmt.Errorf("qty gak boleh 0 atau dibawah nol")
	}

	return &Order{
		ProductName: name,
		Quantity:    qty,
	}, nil
}
