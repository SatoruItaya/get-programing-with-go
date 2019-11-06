package main

import (
	"fmt"
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
				fmt.Print(" ")
			}
		}
		fmt.Println()
		h++
	}
}

func main() {
	universe := NewUniverse()
	universe.Show()
}
