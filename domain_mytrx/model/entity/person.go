package entity

import "fmt"

type Person struct {
	ID   string `bson:"_id"`
	Name string
	Age  int
}

func NewPerson(id, name string, age int) (*Person, error) {

	if age <= 0 {
		return nil, fmt.Errorf("umur gak boleh 0 atau dibawah nol")
	}

	//return &Person{
	//	ID:   id,
	//	Name: name,
	//	Age:  age,
	//}, nil

	return nil, fmt.Errorf("Error spot for NewPerson")

}
