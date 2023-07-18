package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	pos := [8]int{0, 0, 0, 0, 0, 0, 0}
	trouver(0, pos)
}

func Contient(qn int, rp int, pos [8]int) bool {
	for i := 0; i < qn; i++ {
		t := pos[i]
		if t == rp || t == rp-(qn-i) || t == rp+(qn-i) {
			return false
		}
	}
	return true
}

func trouver(nb int, pos [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos[i] + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < 8; i++ {
			if Contient(nb, i, pos) {
				pos[nb] = i
				trouver(nb+1, pos)
			}
		}
	}
}
