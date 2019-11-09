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

func (u Universe) String() string{
	var b byte
	buf := make([]byte, 0, (width+1)*height)
	
	for h:= 0; h < height; h++{
		for w:= 0; w < width; w++ {
			if u[h][w] {
				b = '*'
			}else{
				b = ' '
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (u Universe) Show(){
	fmt.Print("\x0c", u.String())
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

func (u Universe) Set(x, y int, b bool){
	u[x][y] = b
}

func Step(a, b Universe){
	for x := 0; x < height; x++{
		for y := 0; y < width; y++{
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main(){
	a,b := NewUniverse(), NewUniverse()
	a.Seed()
	
	for i := 0; i <= 300; i++{
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a,b = b,a
	}
	
}