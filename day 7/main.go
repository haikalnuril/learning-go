// go routine

package main

import (
	"fmt"
)

// func say (s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main () {
// 	go say("world")
// 	say("hello")
// }


// func sendMessage(ch chan string) {
// 	time.Sleep(2 * time.Second) // Simulate some work
// 	ch <- "Hello from the goroutine!" // Send data into the channel
// 	close(ch) // Close the channel
// }

// func main() {
// 	ch := make(chan string) // Create a channel

// 	go sendMessage(ch) // Start a goroutine

// 	fmt.Println("Waiting for message...")
// 	msg := <-ch // Receive data from the channel
// 	fmt.Println("Message received:", msg)
// }

//learning about select and default in channel. so basicly it just like switch case in other language

func main() {
	ch := make(chan string, 1)

	ch <- "Hello!"
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No message received.")
	}
}
