package entity

import "fmt"

type Person struct {
	Nama string
	Umur int
}

func NewPerson(nama string, umur int) (*Person, error) {

	if umur <= 0 {
		return nil, fmt.Errorf("umur gak boleh 0 atau dibawah nol")
	}

	return &Person{
		Nama: nama,
		Umur: umur,
	}, nil

}
