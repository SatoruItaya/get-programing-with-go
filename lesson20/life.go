package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe{
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show(){
	h := 0
	for h < height{
		for w:= 0; w < width; w++ {
			if u[h][w] {
				fmt.Print("*")
			}else{
				fmt.Print("-")
			}
		}
		fmt.Println()
		h++
	}
}

func (u Universe) Seed(){
	h := 0
	for h < height{
		for w:= 0; w < width; w++ {
			num := rand.Intn(4)
			if num == 0 {
				u[h][w] = true
			}else{
				u[h][w] = false
			}
		}
		h++
	}
}

func (u Universe) Alive(x, y int) bool{
	x = (x + height) % height
	y = (y +  width) % width
	return u[x][y]
}

func (u Universe) Neighbors(x, y int) int{
	count := 0
	for s := -1; s <= 1; s++ {
		for t := -1; t <= 1; t ++ {
			if !(s == 0 && t == 0) && u.Alive(x+s, y+t){
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool{
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func Step(a, b Universe){
	a.Seed()
	for {
		a, b = b, a
		a.Show()
		
		s, t := 0, 0
		for s < height{
			for t < width{
				b.Next(s, t)
			}
		}
		
		time.Sleep(3 * time.Second)
		fmt.Print("\x0c")
	}
}

func main(){
	universe := NewUniverse()
	Step(universe, universe)
}