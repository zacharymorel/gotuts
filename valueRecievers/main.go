package main

// In GOLOANG, there are no classes.

import (
	"fmt"
)

const ussixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gasPedal       uint16 // min 0 max 65535
	breakPedal     uint16
	streelingWheel int16 // -32k +32k
	topSpeedKMH    float64
}

//value reciever
func (c car) khm() float64 {
	// THIS IS A METHOD ON A STRUCT
	return float64(c.gasPedal) * (c.topSpeedKMH / ussixteenbitmax)
}

func (c car) mph() float64 {
	// THIS IS A METHOD ON A STRUCT
	return float64(c.gasPedal) * (c.topSpeedKMH / ussixteenbitmax / kmh_multiple)
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
}
