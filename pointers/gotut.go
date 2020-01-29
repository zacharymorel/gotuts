package main

import "fmt"

// https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and-in-golang
func main() {
	x := 15
	a := &x

	// a := &x
	// is same as
	// var a *int = &x

	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println("-----------------")
	*a = 5
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println("-----------------")
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)

	// & means that is the memory address of w/e you pass
	// * is the way we dereference an reference

	var b = 5
	var p = &b // p holds variable a's memory address
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)

	// Let's change a value (using the initial variable or the pointer)
	*p = 3 // using pointer
	b = 3  // using initial var

	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)
}
