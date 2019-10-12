package main

import (
	"fmt"
	"math/rand"
	
)

func main() {

	fmt.Println("Spaceline        Days Trip type    Price")
	fmt.Println("===============================================")

	distance := 62100000
	spaceline := ""
	tirpType := ""

	for i := 0; i <= 10; i++ {
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		velocity := rand.Intn(15) + 16
		days := distance / velocity / 60 / 60 / 24
		price := velocity + 20

		if rand.Intn(2) == 0 {
			tirpType = "One-way"
		} else {
			tirpType = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-16v %-3v  %-12v $%3v\n", spaceline, days, tirpType, price)
	}
}
