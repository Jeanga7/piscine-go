package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for i := 1; i <= y; i++ { //(i pour les colonnes)
		for j := 1; j <= x; j++ { //(j pour les lignes)
			if ((i == 1 && j == 1) || (i == 1 && j == x)) || ((i == y && j == 1) || (i == y && j == x)) { //(1ere case || derniere case -- 1ere ligne)
				z01.PrintRune('o') // || (1ere case || derniere case -- derniere ligne)
			} else if (i == 1 || i == y) && !(j == 1 || j == x) {
				z01.PrintRune('-') //(les 1ere et derniere ligne sauf les cases des extremités
			} else if (j == 1 || j == x) && !(i == 1 || i == y) {
				z01.PrintRune('|') // les 1ere et derniere colonnes sauf les cases des extremités)
			} else {
				z01.PrintRune(' ') // pour les espaces au milieu du rectangle
			}
		}
		z01.PrintRune('\n')
	}
}
