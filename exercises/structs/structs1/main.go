// structs1
// Make me compile!
//
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func main() {
	person := newPerson("John", 33)
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
