// KEY: VAL STORAGE WITHIN GOLANG

package main

import "fmt"

func main() {
	grades := make(map[string]float32)
	// map is a reference type UNLESS YOU USE go's MAKE()

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 75
	fmt.Println(grades)
	fmt.Println("-----------------------------------")

	TimsGrade1 := grades["Timmy"]
	fmt.Println(TimsGrade1)
	TimsGrade1 = 33
	fmt.Println(TimsGrade1)
	TimsGrade2 := grades["Timmy"]
	fmt.Println(TimsGrade2)
	fmt.Println("-----------------------------------")

	delete(grades, "Timmy")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ": ", v)
	}
}
