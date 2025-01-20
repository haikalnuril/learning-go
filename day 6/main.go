package main

import ("fmt", "time")


//stringer example

type Car struct {
	Type string
	Color string
}

func (c Car) String() string {
	return fmt.Sprintf("My Car type is %s and the color is %s", c.Type, c.Color)
}


func main() {
	car := Car{Type: "SUV", Color: "Red"}
	fmt.Println(car)
}