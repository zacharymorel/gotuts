package main

import "fmt"

// func add(x float64, y float64) float64 {} SHORT HAND BELOW
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	// MULTI RETURN
	return a, b
}

func main() {
	// var num1 float64 = 1
	// var num1 float64 = 1
	num1, num2 := 1.2, 2.1
	// ^ ON COMPILE TIME, GO WILL FIGURE OUT WHAT THE VAR TYPES NEED TO BE, YET THEY CANNOT CHANGE ONCE COMPILED.

	fmt.Println("ADDING: ", add(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
}
