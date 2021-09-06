package main

import (
	"fmt"
	"sync"
)

// func sayHello() {
// 	fmt.Println("Hello")
// }

var wg = sync.WaitGroup{}

func main() {
	msg1 := "hello"
	wg.Add(1)
	go func() {
		fmt.Println("Message 1:", msg1)
		wg.Done()
	}()
	msg1 = "goodbye"
	wg.Wait()

	msg2 := "hello"
	wg.Add(1)
	go func(m string) {
		fmt.Println("Message 2:", m)
		wg.Done()
	}(msg2)
	msg2 = "goodbye"
	wg.Wait()
}
