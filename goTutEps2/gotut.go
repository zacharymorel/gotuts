package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to GO!")
	sqrt()
	generateRandNum()
}

func sqrt() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func generateRandNum() {
	fmt.Println("A Number from 1-100", rand.Intn(100))
}
