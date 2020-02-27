package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done() // ONCE THIS GO ROUTINE CODE IS FINISHED, WE TILL THE WAITGROUP THT THIS GO ROUTINE IS DONE
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10) // BUFFER FOR 10 ITEMS
	for i := 0; i < 10; i++ {
		wg.Add(1) // ADDED VALUES TO WAIT GROUP
		go foo(fooVal, i)
	}

	wg.Wait() // BEFORE WE CLOSE CHANNEL, WE WANT ALL OF THE VALUES TO BE PASSEDD THROUGH THE CHANNEL AND THE WG COUNT IS DOWN TO 0
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

	// Basic example
	// go foo(fooVal, 5)
	// go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal

	// fmt.Println(v1)
	// fmt.Println(v2)
}
