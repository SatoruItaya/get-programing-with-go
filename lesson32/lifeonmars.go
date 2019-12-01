package main

import (
	"fmt"
)

type MarsGrid struct{
}

func (g *MarsGrid) Occupy(p iamge.Point) *Occupier{
}

type Occupier struct{
}

func (g *Occupier) MoveTo(p image.Point) bool{
}

func main() {
}
