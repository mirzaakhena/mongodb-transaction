package entity

import "fmt"

type Order struct {
	ID          string `bson:"_id"`
	ProductName string
	Quantity    int
}

func NewOrder(id, name string, qty int) (*Order, error) {

	if qty <= 0 {
		return nil, fmt.Errorf("qty gak boleh 0 atau dibawah nol")
	}

	return &Order{
		ID:          id,
		ProductName: name,
		Quantity:    qty,
	}, nil
}
