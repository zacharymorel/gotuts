package main

import "fmt"

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
}
