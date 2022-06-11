package main

import (
	"fmt"
	"go-custom-data-types/src/organisation"
	"log"
)

func main() {
	p := organisation.NewPerson("John", "Doe")

	err := p.SetTwitterHandler("@doe.joe")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.ID())
	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler())
}
