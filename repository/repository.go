package repository

import "errors"

// Person
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func all() []Person {
	return []Person{
		Person{"b1", "Horst", 42},
		Person{"c2", "Frank", 28},
		Person{"d1", "Anton", 31},
		Person{"e3", "Steve", 34},
	}
}

// GetPerson returns a Person
func GetPerson(id string) (*Person, error) {
	for _, p := range all() {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Person not found")
}

// GetPeople returns all Persons
func GetPeople() []Person {
	return all()
}
