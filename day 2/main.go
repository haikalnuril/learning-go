package main

import (	
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		} else{
			fmt.Println(i)
		}
	}
	fmt.Println("Kapan hari Sabtu?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Sekarang.")
	case today + 1:
		fmt.Println("Besok.")
	case today + 2:
		fmt.Println("Dua hari lagi.")
	default:
		fmt.Println("Masih jauh.")
	}
}