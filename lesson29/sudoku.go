package main

import (
	"fmt"
	"err"
)

func main() {
	fmt.Println("Hello, playground")
}

const (
	rows, columns = 9, 9
	empty = 0
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit = errors.New("invalid digit")
	ErrInRow = errors.New("digit already present in this row")
	ErrInColumn = errors.New("digit already present in this column")
	ErrInRegion = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

func NewSudoku(digits [rows][columns]int8) *Grid {
}

func (g *Grid) Set(row, column int, digit int8) error {
}

func inBounds(row, column int) bool{
	if row < 0 || column < 0 || row >= rows || column >= columns {
		return false
	}
	return true
}

func validDigit(digit init8){
	return 1 <= digit && digit <= 9
}