package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune(120)
	z01.PrintRune(32)
	z01.PrintRune(61)
	z01.PrintRune(32)
	z01.PrintRune(52)
	z01.PrintRune(50)
	z01.PrintRune(44)
	z01.PrintRune(32)
	z01.PrintRune(121)
	z01.PrintRune(32)
	z01.PrintRune(61)
	z01.PrintRune(32)
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune(10)
}
