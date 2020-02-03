package main

// In GOLOANG, there are no classes.

import (
	"fmt"
)

const usSixteenBitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal       uint16 // min 0 max 65535
	breakPedal     uint16
	streelingWheel int16 // -32k +32k
	topSpeedKMH    float64
}

//DIFF BETWEEN VAL AND POINTER RECEIVERS
// val receiver make a copy of the strcut they are on
// pointer receiver return the values at the address of the pointer of the strcut they are on

//value receiver
func (c car) khm() float64 {
	// THIS IS A METHOD ON A STRUCT
	fmt.Println("cBefore: ", c)
	return float64(c.gasPedal) * (c.topSpeedKMH / usSixteenBitMax)
}

func (c car) mph() float64 {
	// THIS IS A METHOD ON A STRUCT
	return float64(c.gasPedal) * (c.topSpeedKMH / usSixteenBitMax / kmhMultiple)
}

// pointer receiver
func (c *car) newTopSpeed(newSpeed float64) { // pointer within method signature is because we need to tell the struct that this is a pointer receiver and we need to pass in the actually vals at address
	fmt.Println("cBefore: ", *c)
	c.topSpeedKMH = newSpeed
	fmt.Println("cAfter: ", *c)
}

// FUNCTION THAT DOES THE SAME THING AS VAL AND POINTER RECEIVERS
func newestTopSpeed(c car, speed float64) car {
	c.topSpeedKMH = speed
	return c
}

func main() {
	aCar := car{
		gasPedal:       65000,
		breakPedal:     0,
		streelingWheel: 12561,
		topSpeedKMH:    225.20,
	}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.khm())
	fmt.Println(aCar.mph())
	fmt.Println("--------------------------------------")

	aCar.newTopSpeed(500)
	fmt.Println(aCar.khm())
	fmt.Println(aCar.mph())
	// fmt.Println("--------------------------------------")

	// aCar = newestTopSpeed(aCar, 500)
	// fmt.Println(aCar.khm())
	// fmt.Println(aCar.mph())
}
