package organisation

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type TwitterHandler string
type SocialSecurityNumber string

func (ssn SocialSecurityNumber) ID() string {
	return string(ssn)
}

func NewSocialSecurityNumber(value string) Identifiable {
	return SocialSecurityNumber(value)
}

type Name struct {
	first string
	last  string
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstName string, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
}

// Struct methods

func (th TwitterHandler) RediredctUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

//func (p *Person) ID() string {
//	return "hello-world"
// }

// Pointer-based receiver

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
