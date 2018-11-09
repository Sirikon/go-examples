package main

import "fmt"

// Person .
type Person struct {
	ID      int
	Name    string
	Surname string
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(5 * 3)

	person := Person{ID: 20, Name: "Carlos", Surname: "Fernández"}

	fmt.Println(person)
}
