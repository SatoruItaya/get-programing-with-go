package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-16v %v   %-12v  %3v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================================")

	for i := 0; i <= 10; i++ {
		switch rand.Intn(3) {
		case 0:
			fmt.Printf("%-16v ", "Space Adventures")
		case 1:
			fmt.Printf("%-16v ", "SpaceX")
		case 2:
			fmt.Printf("%-16v ", "Virgin Galactic")
		}
		days := rand.Intn(15) + 16
		price := days + 20
		fmt.Printf("%v     ", days)

		if rand.Intn(2) == 0 {
			fmt.Printf("%-13v $%4v\n", "One-way", price)
		} else {
			fmt.Printf("%-13v $%4v\n", "Round-trip", price*2)
		}
	}
}
