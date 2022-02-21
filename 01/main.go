package main

import (
	"fmt"
)

type Person struct {
	lastName   string
	firstName  string
	patronymic string
}

type Action struct {
	Person
}

func (p Person) getName() {
	fmt.Printf("Person's full name %s %s %s", p.lastName, p.firstName, p.patronymic)
}

func (p *Person) setName(lastName string, firstName string, patronymic string) {
	p.lastName = lastName
	p.firstName = firstName
	p.patronymic = patronymic
}

func main() {
	action := new(Action)
	action.setName("Sargisyan", "Samvel", "Manvelovich")
	action.getName()
}
