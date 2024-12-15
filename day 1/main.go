package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var i int = 42
	j :=207
	const pi float32 = 3.14
	fmt.Println("Hello, World!", add(1, 2), i, j, pi)
}