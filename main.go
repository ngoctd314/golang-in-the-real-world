package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Saver interface {
	Save(data []byte) error
}

func SavePerson(person *Person, saver Saver) error {
	// validate input
	err := person.validate()
	if err != nil {
		return err
	}

	// encode person to bytes
	bytes, err := person.encode()
	if err != nil {
		return err
	}

	// save the person and return the result
	return saver.Save(bytes)
}

type Person struct {
	Name  string
	Phone string
}

func (p *Person) validate() error {
	if p.Name == "" {
		return errors.New("name missing")
	}
	if p.Phone == "" {
		return errors.New("phone missing")
	}

	return nil
}
func (p *Person) encode() ([]byte, error) {
	return json.Marshal(p)
}

func LoadPerson(ID int, decodePerson func(data []byte) *Person) (*Person, error) {
	// validate the input
	if ID <= 0 {
		return nil, fmt.Errorf("invalid ID '%d' supplied", ID)
	}

	// load from storage
	bytes, err := loadPerson(ID)
	if err != nil {
		return nil, err
	}

	// decode bytes and return

	return decodePerson(bytes), nil
}
func loadPerson(ID int) ([]byte, error) {
	return nil, errors.New("not implementd")
}

func main() {

}
