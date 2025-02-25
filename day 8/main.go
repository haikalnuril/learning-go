package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// func sendData(ch chan<- int) {
// 	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

// 	for i := 0; true; i++ {
// 		ch <- i
// 		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
// 	}
// }

// func retreiveData(ch <-chan int) {
// loop:
// 	for {
// 		select {
// 		case data := <-ch:
// 			fmt.Print(`receive data "`, data, `"`, "\n")
// 		case <-time.After(time.Second * 5):
// 			fmt.Println("timeout. no activities under 5 seconds")
// 			break loop
// 		}
// 	}
// }

func testDefer() {
	defer fmt.Println("halo")
	// os.Exit(1)
	fmt.Println("Selamat datang")
}

func errorTest() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
    if r := recover(); r != nil {
        fmt.Println("Error occured", r)
    } else {
        fmt.Println("Application running perfectly")
    }
}

func panicTest() {
	var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
    }
}

func main() {
	// runtime.GOMAXPROCS(2)

	// var messages = make(chan int)

	// go sendData(messages)
	// retreiveData(messages)

	
	defer catch()

	testDefer()
	errorTest()
	panicTest()
}
