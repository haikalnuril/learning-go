package main

import "fmt"

func main() {
	// how to create array
	var a [6]int

	// how to assign value to array
	array := [6]int{1, 2, 3, 4, 5, 6}

	//how to slice array
	slice := array[1:4]

	// how to create array w/o specifying the length
	r := []bool{true, false, true, true, false, true}

	// how to append to slice
	s:= []int{}
	s = append(s, 1)

	// Slice input
    items := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

    // Map untuk menghitung frekuensi
    frequency := make(map[string]int)

    // Iterasi slice dan update map
    for _, item := range items {
        frequency[item]++
    }

    // Menampilkan hasil
    fmt.Println("Frekuensi elemen:")
    for key, count := range frequency {
        fmt.Printf("%s: %d\n", key, count)
    }

	fmt.Println(r)
	fmt.Println(s)

	fmt.Println(a)
	fmt.Println(array)
	fmt.Println(slice)
}