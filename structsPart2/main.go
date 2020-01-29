package main

import "fmt"

type person struct {
	name string
	age  int
}

// https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and-in-golang
func NewPerson(name string) *person { // <= *person within func signatures says, "this func returns a pointer of type person (struct)"
	// This is a constructor function for person
	p := person{name: name}
	p.age = 42
	return &p // since we return the pointer, that means this variable (through it's memory address) can now be access outside of the func scope where it was defined.
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	f := NewPerson("Jon")
	fmt.Println(*f)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
