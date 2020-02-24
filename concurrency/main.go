package main

import (
	"fmt"
	"sync"
	"time"
)

// DEFER STATEMENT SAYS, Defer running this block of code UNTIL surrounding Scope is completed/ returned or ERRORED OUT
// WE CAN USE DEFER TO RECOVER AFTER AN ERROR // PANIC // RECOVER
// GENERALLY YOU PUT A RECOVER FUNC WITHIN DEFER

var wg sync.WaitGroup

func cleanUp() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("RECOVERED IN CLEANUP ", r)
	}
}

func say(s string) {
	// defer wg.Done()
	defer cleanUp()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("oh dear, a 2")
		}
	}
}

func main() {
	wg.Add(1)
	go say("HEY")
	wg.Add(1)
	go say("THERE")
	wg.Wait()

	// wg.Add(1)
	// go say("Hi")
	// wg.Wait()
}

// DEFER
// func foo() {
// 	defer fmt.Println("DONE")
// 	defer fmt.Println("Are we Done?") // Defer statements work in first in, last out
// 	fmt.Println("DOING THINGS")
// }

// func main() {
// 	foo()
// }
