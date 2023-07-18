package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	pos := []int{0, 0, 0, 0, 0, 0, 0, 0}
	trouver(0, pos)
}
func contient(qn int, rp int, pos [8]int) bool {
	for i := 0; i < qn; i++ {
		t := pos[i]
		if t == rp || t == rp-(qn-1) || t == rp+(qn-1) {
			return false
		}
	}
	return true
}
func trouver(nb int, pos [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos(i) + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < 8; i++ {
			if contient(nb, i, pos) {
				pos[nb] = i
				trouver(nb+1, pos)
			}
		}
	}
}
