package main 

import "fmt"

type Drivable interface {
	Drive()
}

type Car struct {
	Name string
}

func (c Car) Drive() {
	fmt.Printf(" %s Car is On The Way\n", c.Name)
}

func main() {
	car := Car{Name: "Ferrari"}
	car.Drive()
}