package main

import (
	"fmt"
	"math/rand"
	"time"
)

type gopher struct {
	name string
}

func (g gopher) string() string{
	return g.name
}

func (g gopher) move() string {
	switch  rand.Intn(2){
	case 0:
		return "scurries along the ground"
	default:
		return "burrows in the sand"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(2){
	case 0:
		return "carrot"
	default:
		return "lettuce"
	}
}

type animal interface {
	string() string
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2){
	case 1:
		fmt.Printf("%v %v.\n", a.string(), a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a.string(), a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())
	
	animals := []animal{
		gopher{name: "GO gopher"},
	}
	
	var sol, hour int
	
	for{
		fmt.Printf("%2d:00", hour)
		if hour < sunrise || hour >= sunset{
			fmt.Println("The animals are sleeping")
		}else{
			i := rand.Intn(len(animals))
			step(animals[i])
		}
		
		time.Sleep(500 * time.Millisecond)
		
		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3{
				break
			}
		}
	}
}
