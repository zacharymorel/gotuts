package main

// In GOLOANG, there are no classes.

import (
	"fmt"
)

type car struct {
	gasPedal       uint16 // min 0 max 65535
	breakPedal     uint16
	streelingWheel int16 // -32k +32k
	topSpeedKMH    float64
}

func main() {
	aCar := car{
		gasPedal:       22341,
		breakPedal:     0,
		streelingWheel: 12561,
		topSpeedKMH:    225.20,
	}

	fmt.Println(aCar.gasPedal)
}
