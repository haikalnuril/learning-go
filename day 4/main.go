package main

import "fmt"

type Car struct {
	Color string
	Name  string
	Type string
	Speed int
}

func (c Car) Menyala() {
	fmt.Println(fmt.Sprintf("Mobil %s menyala", c.Name))
}

// ilmu mahal
// go have a 2 type of receiver -> pointer receiver and value receiver
// pointer receiver -> if you want to change the value of the struct permanently
// value receiver -> if you want to change the value of the struct temporarily and in that function scope only

// Method with a pointer receiver to modify the original car
func (c *Car) Accelerate(change int) {
	c.Speed += change
}

// Function (not a method) that requires a pointer to modify the car
func AccelerateFunc(c *Car, change int) {
	c.Speed += change
}

func main() {
	car := Car{Color: "Red", Name: "Ferrari", Type: "Sport", Speed: 200}

	car.Menyala()

	// Example 1: Using a method with a pointer receiver
	car1 := Car{Speed: 50}
	car1.Accelerate(20) // Go automatically converts "car1" to "&car1"
	fmt.Println("Car1 Speed:", car1.Speed) // Output: 70

	// Example 2: Using a function that takes a pointer
	car2 := Car{Speed: 40}
	// AccelerateFunc(car2, 30) // Compile error: requires &car2
	AccelerateFunc(&car2, 30) // Works because we pass the address explicitly
	fmt.Println("Car2 Speed:", car2.Speed) // Output: 70
}