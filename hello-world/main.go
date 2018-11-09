package main

import "fmt"

// Person .
type Person struct {
	ID      int
	Name    string
	Surname string
}

// GetFullName .
func (p *Person) GetFullName() string {
	return p.Name + " " + p.Surname
}

// SetName .
func (p *Person) SetName(name string) {
	p.Name = name
}

// SetPersonName .
func SetPersonName(person *Person, name string) {
	person.Name = name
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(5 * 3)

	person := &Person{ID: 20, Name: "Carlos", Surname: "Fernández"}

	person.SetName("Peter")
	SetPersonName(person, "Johnny")

	fmt.Println(person)
	fmt.Println(person.GetFullName())
}
