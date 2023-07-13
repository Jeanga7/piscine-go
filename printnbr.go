package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	z01.PrintRune('n')
}
func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
