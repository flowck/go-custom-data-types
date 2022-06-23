package main

import (
	"fmt"
	"go-custom-data-types/src/organisation"
	"log"
)

func main() {
	p := organisation.NewPerson("John", "Doe", organisation.NewSocialSecurityNumber("123-456-789"))

	err := p.SetTwitterHandler("@doe.joe")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RediredctUrl())
	fmt.Println(p.ID())

	person1 := &Name{first: "", last: ""}
	// person2 := Name{first: "John", last: "Doe"}
	// portfolio := map[Name][]organisation.Person{}

	if person1 == nil {
		fmt.Println("We match")
	}
}

type Name struct {
	first string
	last  string
}
